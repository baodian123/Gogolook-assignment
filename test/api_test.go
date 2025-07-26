package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/baodian123/Gogolook-assignment/internal/api/controller"
	"github.com/baodian123/Gogolook-assignment/internal/application/services"
	"github.com/baodian123/Gogolook-assignment/internal/infrastructure/repository"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	repo := repository.NewInMemoryTaskRepository()
	svc := services.NewTaskService(repo)
	ctrl := controller.NewTaskController(svc)
	r := gin.Default()
	ctrl.RegisterRoutes(r)
	return r
}

func TestCreateTask_Success(t *testing.T) {
	r := setupRouter()

	payload := map[string]interface{}{
		"name":   "Test Task",
		"status": 1,
	}

	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	var resp map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}

	id, ok := resp["id"].(string)

	if !ok || id == "" {
		t.Fatalf("expected non-empty id")
	}
}

func TestCreateTask_InvalidStatus(t *testing.T) {
	r := setupRouter()

	payload := map[string]interface{}{
		"name":   "Test Task",
		"status": 99,
	}

	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}

	var resp map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}

	if resp["error"] != "unknown task status" {
		t.Errorf("expected error 'unknown task status', got %v", resp["error"])
	}
}

func TestGetTaskList_Empty(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp []interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}

	if len(resp) != 0 {
		t.Errorf("expected empty list, got %v", resp)
	}
}

func TestTask_CRUD_Flow(t *testing.T) {
	r := setupRouter()

	// Create
	payload := map[string]interface{}{
		"name":   "Task1",
		"status": 1,
	}

	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	var createResp map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &createResp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}

	id, ok := createResp["id"].(string)

	if !ok || id == "" {
		t.Fatalf("expected non-empty id")
	}

	// Get List
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var listResp []map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &listResp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}

	if len(listResp) != 1 || listResp[0]["id"] != id {
		t.Errorf("expected 1 task with correct id, got %v", listResp)
	}

	// Update
	updatePayload := map[string]interface{}{
		"name":   "Task1 Updated",
		"status": 0,
	}

	updateBody, _ := json.Marshal(updatePayload)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/tasks/"+id, bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var updateResp map[string]interface{}

	if err := json.Unmarshal(w.Body.Bytes(), &updateResp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}

	if updateResp["name"] != "Task1 Updated" || int(updateResp["status"].(float64)) != 0 {
		t.Errorf("update response mismatch: %v", updateResp)
	}

	// Delete
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/tasks/"+id, nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", w.Code)
	}

	// Get List after delete
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if err := json.Unmarshal(w.Body.Bytes(), &listResp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}

	if len(listResp) != 0 {
		t.Errorf("expected empty list after delete, got %v", listResp)
	}
}
