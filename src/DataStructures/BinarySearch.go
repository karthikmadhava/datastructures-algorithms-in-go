package DataStructures

func binarySearch(arr []int, target int) int {
	// Set low and high boundaries
	low, high := 0, len(arr)-1

	for low <= high { // Continue search as long as low <= high
		// Calculate middle index
		mid := low + (high-low)/2
		// Check if element is found
		if target == arr[mid] {
			return mid
		}

		if arr[mid] < target {
			// If target is greater, focus on the right half
			low = mid + 1
		} else {
			// If target is less, focus on the left half
			high = mid - 1
		}
	}

	// If element not found, return -1
	return -1
}
