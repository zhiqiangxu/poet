package shared

import (
	"github.com/spacemeshos/sha256-simd" // simd optimized sha256 computation
	"hash"
)

// HashFunc implementation
type sha256Hash struct {
	x          []byte // arbitrary binary data
	hash       hash.Hash
	iters      int
	emptySlice []byte
}

// Returns a new HashFunc Hx() for commitment X
func NewHashFunc(x []byte) HashFunc {

	// todo: pick iter value form params
	iters := 50

	return &sha256Hash{x: x, hash: sha256.New(), iters: iters}
}

// Hash implements Hx()
func (h *sha256Hash) HashTemp(data ...[]byte) []byte {
	h.hash.Reset()
	h.hash.Write(h.x)
	for _, d := range data {
		_, _ = h.hash.Write(d)
	}

	return h.hash.Sum(h.emptySlice)

}

// Multiple iterations hash using client provided iters
func (h *sha256Hash) Hash(data ...[]byte) []byte {

	h.hash.Reset()

	// first, hash x
	h.hash.Write(h.x)

	// hash all user provided data
	for _, d := range data {
		_, _ = h.hash.Write(d)
	}

	digest := h.hash.Sum([]byte{})

	// perform iter hashes of x and user data
	for i := 0; i < h.iters; i++ {
		h.hash.Reset()
		h.hash.Write(h.x)
		h.hash.Write(digest)
		digest = h.hash.Sum(h.emptySlice)
	}

	return digest
}

func (h *sha256Hash) HashSingle(data []byte) []byte {
	h.hash.Reset()
	h.hash.Write(h.x)
	h.hash.Write(data)
	return h.hash.Sum(h.emptySlice)
}
