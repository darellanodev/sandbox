package main

type GameObject interface {
	Reset()
	MoveToX(float64)
	FlipX()
	Alpha(float64)
	Rotate(float64)
	White()
}

func ApplyMovement[T1 GameObject](obj T1, posX float64) T1 {
	obj.Reset()
	obj.MoveToX(posX)

	return obj
}

func ApplyMovementFlipX[T1 GameObject](obj T1, posX float64) T1 {
	obj.Reset()
	obj.FlipX() // the order is important. If you want to flip the object do it before moving it.
	obj.MoveToX(posX)

	return obj
}

func ApplyMovementAlpha[T1 GameObject](obj T1, posX float64, alpha float64) T1 {
	obj.Reset()
	obj.MoveToX(posX)
	obj.Alpha(alpha)

	return obj
}

func ApplyMovementRotationAlpha[T1 GameObject](obj T1, posX float64, alpha float64, degrees float64) T1 {
	obj.Reset()
	obj.Rotate(degrees) // the order is important. If you want to rotate the object do it before moving it.
	obj.MoveToX(posX)
	obj.Alpha(alpha)

	return obj
}

func ApplyMovementWhite[T1 GameObject](obj T1, posX float64) T1 {
	obj.Reset()
	obj.White()
	obj.MoveToX(posX)

	return obj
}