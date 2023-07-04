package twitter_uploader

import (
	"testing"
)

func TestCreateBatches(t *testing.T) {
	batches := createBatches(208175)
	if len(batches) != 1 {
		t.Errorf("Expected 1 batches, got %v", len(batches))
	}

	batches = createBatches(20817500)
	if len(batches) != 4 {
		t.Errorf("Expected 10 batches, got %v", len(batches))
	}

	if batches[0] != BatchSize {
		t.Errorf("Expected first batch to be %v, got %v", BatchSize, batches[0])
	}

	if batches[1] != BatchSize {
		t.Errorf("Expected second batch to be %v, got %v", BatchSize, batches[1])
	}

	if batches[2] != BatchSize {
		t.Errorf("Expected third batch to be %v, got %v", BatchSize, batches[2])
	}

	if batches[3] != 5088860 {
		t.Errorf("Expected fourth batch to be %v, got %v", 17500, batches[3])
	}
}
