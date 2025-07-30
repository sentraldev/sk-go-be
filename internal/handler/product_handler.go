package handler

import (
	"fmt"
	"log"
	"net/http"
	"sk-go-be/internal/model"
	"sk-go-be/internal/service"
	"sk-go-be/internal/utils"
	types "sk-go-be/internal/utils/type"

	"firebase.google.com/go/v4/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	ProductService service.ProductService
	StorageClient  storage.Client // Assuming you have a StorageClient in utils package
}

func NewProductHandler(productService service.ProductService, storageClient storage.Client) *ProductHandler {
	return &ProductHandler{ProductService: productService, StorageClient: storageClient}
}

func (h *ProductHandler) GetProductByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	product, err := h.ProductService.GetProductByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for j, imageURL := range product.ImageURLs {
		url := utils.GetFirebasePublicUrl(imageURL)
		if err != nil {
			log.Printf("Error generating presigned URL for product %s image %d: %v", product.UUID, j, err)
			continue // Skip this image if there's an error
		}
		product.ImageURLs[j] = url // Update the image URL with the presigned URL
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate presigned URLs for images
	for i, product := range products {
		for j, imageURL := range product.ImageURLs {
			url := utils.GetFirebasePublicUrl(imageURL)
			if err != nil {
				log.Printf("Error generating presigned URL for product %s image %d: %v", product.UUID, j, err)
				continue // Skip this image if there's an error
			}
			products[i].ImageURLs[j] = url // Update the image URL with the presigned URL
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// ADMIN

type ProductDTO struct {
	Name         string  `form:"name" binding:"required"`
	Description  string  `form:"description" binding:"required"`
	Price        float64 `form:"price" binding:"required"`
	CategoryUUID string  `form:"category_uuid" binding:"required"`
	BrandUUID    string  `form:"brand_uuid" binding:"required"`
	Highlight    string  `form:"highlight" binding:"required"`
}

func (h *ProductHandler) CreateNewProduct(c *gin.Context) {
	// Parse multipart form (set max memory limit, e.g., 32MB)
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Bind form fields to ProductDTO
	var product ProductDTO
	if err := c.ShouldBind(&product); err != nil {
		log.Printf("Error binding form data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid form data: %v", err)})
		return
	}

	log.Printf("Received product: %+v", product)

	// Get all files under "photos[]"
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Error accessing multipart form: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to access multipart form"})
		return
	}

	files := form.File["photos[]"]
	if len(files) == 0 {
		log.Printf("No photos provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one photo is required"})
		return
	}

	// Validate and upload each file
	var imageUrls []string
	for i, file := range files {
		// Validate file type
		contentType := file.Header.Get("Content-Type")
		if contentType != "image/png" && contentType != "image/jpeg" {
			log.Printf("Invalid file type for file %d: %s", i, contentType)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("File %d: Only PNG and JPEG images are allowed", i)})
			return
		}

		// Validate file size (5MB)
		if file.Size > 5<<20 {
			log.Printf("File %d too large: %d bytes", i, file.Size)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("File %d: Image file too large (max 5MB)", i)})
			return
		}

		// Upload file to Firebase
		imageUrl, err := utils.UploadFileToFirebase(&h.StorageClient, file, "products")
		if err != nil {
			log.Printf("Error uploading file %d: %v", i, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload file %d", i)})
			return
		}
		imageUrls = append(imageUrls, imageUrl)
	}

	log.Printf("Uploaded image URLs: %v", imageUrls)

	// Create new product
	var newProduct model.Product
	newProduct.Name = product.Name
	newProduct.Description = product.Description
	newProduct.Price = product.Price
	newProduct.CategoryUUID, _ = uuid.Parse(product.CategoryUUID)
	newProduct.BrandUUID, _ = uuid.Parse(product.BrandUUID)
	newProduct.Highlight = product.Highlight
	newProduct.ImageURLs = types.JSONBArray(imageUrls)

	// Save product to database
	if err := h.ProductService.CreateNewProduct(&newProduct); err != nil {
		log.Printf("Error creating product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create product: %v", err)})
		return
	}

	// Respond with created product details
	c.JSON(http.StatusCreated, gin.H{
		"message":    "Product created successfully",
		"product":    product,
		"image_urls": imageUrls,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paramUuid := c.Param("uuid") // Assuming UUID is passed as a URL parameter
	formattedUUID, err := uuid.Parse(paramUuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.UUID = formattedUUID

	if err := h.ProductService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Optional: If you want to support deletion of products
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	paramUuid := c.Param("uuid") // Assuming UUID is passed as a URL parameter
	if err := h.ProductService.DeleteProduct(paramUuid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil) // No content response for successful deletion
}
