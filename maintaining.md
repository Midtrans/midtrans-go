> Warning: This note is for developer/maintainer of this package only

## Updating Package
- Make your changes
- Update `libraryVersion` value on `midtrans.go` file
- To install dev dependencies with go module run `go mod tidy` on repo folder
- To run test, run 
  - `go test github.com/midtrans/midtrans-go/coreapi`
  - `go test github.com/midtrans/midtrans-go/snap`
  - `go test github.com/midtrans/midtrans-go/iris`

## Release new version
- Commit and push changes to Github master branch
- Create a [Github Release](https://github.com/Midtrans/midtrans-go/releases) with the target version
- Github Release and Master Branch are automatically synced to [Go Packages](https://pkg.go.dev/github.com/midtrans/midtrans-go)