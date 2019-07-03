package domain
import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestNewComputer(t *testing.T){
	comp:=NewComputer()
	assert.NotNil(t,comp)
}