package poseidon

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/iden3/go-iden3-crypto/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/blake2b"
)

func TestBlake2bVersion(t *testing.T) {
	h := blake2b.Sum256([]byte("poseidon_constants"))
	assert.Equal(t, "e57ba154fb2c47811dc1a2369b27e25a44915b4e4ece4eb8ec74850cb78e01b1", hex.EncodeToString(h[:]))
}

func TestPoseidon(t *testing.T) {
	b1 := big.NewInt(int64(1))
	b2 := big.NewInt(int64(2))
	h, err := Hash([]*big.Int{b1, b2})
	assert.Nil(t, err)
	assert.Equal(t, "12242166908188651009877250812424843524687801523336557272219921456462821518061", h.String())

	b3 := big.NewInt(int64(3))
	b4 := big.NewInt(int64(4))
	h, err = Hash([]*big.Int{b3, b4})
	assert.Nil(t, err)
	assert.Equal(t, "17185195740979599334254027721507328033796809509313949281114643312710535000993", h.String())

	msg := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	n := 31
	msgElems := make([]*big.Int, 0, len(msg)/n+1)
	for i := 0; i < len(msg)/n; i++ {
		v := new(big.Int)
		utils.SetBigIntFromLEBytes(v, msg[n*i:n*(i+1)])
		msgElems = append(msgElems, v)
	}
	if len(msg)%n != 0 {
		v := new(big.Int)
		utils.SetBigIntFromLEBytes(v, msg[(len(msg)/n)*n:])
		msgElems = append(msgElems, v)
	}
	hmsg, err := Hash(msgElems)
	assert.Nil(t, err)
	assert.Equal(t, "11821124228916291136371255062457365369197326845706357273715164664419275913793", hmsg.String())

	msg2 := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet.")
	msg2Elems := make([]*big.Int, 0, len(msg2)/n+1)
	for i := 0; i < len(msg2)/n; i++ {
		v := new(big.Int)
		utils.SetBigIntFromLEBytes(v, msg2[n*i:n*(i+1)])
		msg2Elems = append(msg2Elems, v)
	}
	if len(msg2)%n != 0 {
		v := new(big.Int)
		utils.SetBigIntFromLEBytes(v, msg2[(len(msg2)/n)*n:])
		msg2Elems = append(msg2Elems, v)
	}
	hmsg2, err := Hash(msg2Elems)
	assert.Nil(t, err)
	assert.Equal(t, "10747013384255785702102976082726575658403084163954725275481577373644732938016", hmsg2.String())

	hmsg2, err = HashBytes(msg2)
	assert.Nil(t, err)
	assert.Equal(t, "10747013384255785702102976082726575658403084163954725275481577373644732938016", hmsg2.String())
}