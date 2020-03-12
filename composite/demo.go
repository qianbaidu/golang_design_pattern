package main

// region start composite swimmer demo
type ITrainer interface {
	Train()
}

type ISwimmer interface {
	Swim()
}

type CompositeSwimmer struct {
	ITrainer
	ISwimmer
}

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

type Swimmer struct{}

func (b *Swimmer) Swim() {
	println("Swimming")
}

// endregion end composite swimmer demo

// region start tree demo
type Tree struct {
	LeafValue int
	Left      *Tree
	Right     *Tree
}

// endregion end tree demo

func main() {
	swimmer := CompositeSwimmer{
		ITrainer: &Athlete{},
		ISwimmer: &Swimmer{},
	}

	swimmer.Train()
	swimmer.Swim()

	root := Tree{
		LeafValue: 0,
		Left: &Tree{
			LeafValue: 20,
			Left:      nil,
			Right:     nil,
		},
		Right: &Tree{
			LeafValue: 20,
			Left:      nil,
			Right: &Tree{
				LeafValue: 8,
				Left:      nil,
				Right:     nil,
			},
		},
	}

	println(root.Right.Right.LeafValue)

}
