package puppy

func Bark() string {

	return "woof"

}

func Barks() string {

	return "WOOF WOOF!"

}

func BigBark() string {
	return dog.WhenGrownUp(Barks())
}
