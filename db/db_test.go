package db

// func TestCreateUser(t *testing.T) {
// 	tests := []struct {
// 		name 	 string
// 		id       int
// 		username string
// 		balance  int
// 	}{
// 		{
// 			name:     "#1 success",
// 			id:       1,
// 			username: "John",
// 			balance:  400,
// 		},
// 		{
// 			name:     "#2 success",
// 			id:       2,
// 			username: "Barbosa",
// 			balance:  550,
// 		},
// 	}
// 	for _, test := range tests {
// 		if infoUser := UserStorage.U.CreateUser(pkg.User{test.id, test.username, test.balance}); infoUser != pkg.User{test.id,test.name, test.balance} {
// 			t.Fatalf("Expected username or expected balance is not correst!!!!")
// 		}
// 	}
// }

// func TestDeleteUser(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		id   int
// 	}{
// 		{
// 			name: "#1 success",
// 			id:   1,
// 		},
// 		{
// 			name: "#2 success",
// 			id:   2,
// 		},
// 	}
// 	for _, test := range tests {
// 		if infoUser := UserStorage.U.DeleteUser(test.id); infoUser != nil {
// 			t.Fatalf("User is not deleted!!!!")
// 		}
// 	}
// }

// func TestUpdateUser(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		id       int
// 		username string
// 		balance  int
// 	}{
// 		{
// 			name:     "#1 success",
// 			id:       1,
// 			username: "John",
// 			balance:  400,
// 		},
// 		{
// 			name:     "#2 success",
// 			id:       3,
// 			username: "Barbosa",
// 			balance:  550,
// 		},
// 	}
// 	for _, test := range tests {
// 		if infoUser := UserStorage.U.UpdateUser(pkg.User{test.id, test.username, test.balance}); infoUser !=  pkg.User{test.id, test.username, test.balance}{
// 			t.Fatalf("Expected username or expected balance is not correst!!!!")
// 		}
// 	}
// }
