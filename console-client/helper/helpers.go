package helper
import (
	"net/url"
	"fmt"
	"crypto/md5"
	"io"
	"encoding/hex"
	"console-client/domain"
)
func AddParamToURL(u, addfield, value string) string{
	ur, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	q := ur.Query()
	q.Add(addfield, value)
	ur.RawQuery = q.Encode()
	return ur.String()
}

func GetHash(s string) string{
	h := md5.New()
	io.WriteString(h,s)
	return hex.EncodeToString(h.Sum(nil))
}
func PrintPC(pc *domain.PC){
	fmt.Printf("\nComputer with ID: %v",pc.ID)
	fmt.Printf("\n     Inventory number: %v, Vendor: %v",pc.Computer.InventoryNumber,pc.Computer.Vendor)
	fmt.Printf("\n     Specifications: HDD: %v, RAM: %v, Core: %v %v %vGHz",
					pc.Computer.HDDVolume,
					pc.Computer.RAMVolume,
					pc.Computer.Core.CoreVendor, 
					pc.Computer.Core.Model, 
					pc.Computer.Core.Frequency)
	fmt.Printf("\n     Owners infomation: First name: %v, Last name: %v, Office number: %v",
					pc.Computer.Owner.FirstName,
					pc.Computer.Owner.LastName,
					pc.Computer.Owner.RoomNumber)
	fmt.Println("\n_______________________________________________________________")	
}
func PrintUser(user *domain.UserID){
	fmt.Printf("\nUser with ID: %v",user.ID)
	fmt.Printf("\n     Login: %v",user.User.Login)
	fmt.Printf("\n     Rights: %v",user.User.UserRights)
	fmt.Println("\n_______________________________________________________________")	
}