package services

// GetAllPostsService gets all post from the database
func GetAllPostsService() (string, error) {
	return "Getting App Post.", nil
}

// GetPostByIdService retreive single post by it's ID
func GetPostByIdService() (string, error) {
	return "Get Post By Id", nil
}

// CreatePostService create a new post into the database
func CreatePostService() (string, error) {
	return "Create New Post", nil
}

// UpdatePostService update a post in the database
func UpdatePostService() (string, error) {
	return "Update A Post.", nil
}

// DeletePostService delete a post from the database
func DeletePostService() (string, error) {
	return "Delete A Post", nil
}
