package vec

// Make Make
func Make(size int) Vector {
	slice := make([]float64, size)
	return Vector{slice}
}

// Init Init
func Init(nums ...float64) Vector {
	return Vector{nums}
}

// Sub Sub
func Sub(v1, v2 Vector) Vector {
	v := Copy(v1)
	v.Sub(v2)

	return v
}

// Copy Copy
func Copy(from Vector) Vector {
	to := Make(from.Len())
	copy(to.slice, from.slice)

	return to
}

// Slice Slice
func Slice(v Vector, from, to int) Vector {
	return Vector{v.slice[from:to]}
}
