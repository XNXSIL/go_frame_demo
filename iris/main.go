package main

import "github.com/kataras/iris/v12"

type (
	request struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	response struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}
)

func main() {
	app := iris.New()
	app.Handle("PUT", "/users/{id:uuid}", updateUser)
	app.Listen(":8080")
}

func updateUser(ctx iris.Context) {
	id, _ := ctx.Params().Get("id")

	var req request
	if err := ctx.ReadJSON(&req); err != nil {
		// ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	resp := response{
		ID:      id,
		Message: req.Firstname + " updated successfully",
	}
	ctx.JSON(resp)
}
