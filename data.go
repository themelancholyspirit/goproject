package main

type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	IsActive bool
}

var Users = []User{
	{1, "John Doe", "john.doe@example.com", 30, true},
	{2, "Jane Smith", "jane.smith@example.com", 25, true},
	{3, "Michael Johnson", "michael.johnson@example.com", 35, false},
	{4, "Emily Davis", "emily.davis@example.com", 28, true},
	{5, "James Brown", "james.brown@example.com", 32, true},
	{6, "Sarah Wilson", "sarah.wilson@example.com", 27, true},
	{7, "David Lee", "david.lee@example.com", 40, true},
	{8, "Jennifer Taylor", "jennifer.taylor@example.com", 33, true},
	{9, "Brian Clark", "brian.clark@example.com", 29, false},
	{10, "Olivia Rodriguez", "olivia.rodriguez@example.com", 31, true},
}
