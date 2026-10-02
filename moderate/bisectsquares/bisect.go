package bisectsquares

/*

Given two squares on a 2-d plane
find a line that would cut these 2 squares in half
Assume the top and bottom run parallel to x-axis => orthogonal, we dont need slope

A line bisects a square if it passes through its middle

To bisect two squares it has to pass through both middles

*/

type Point struct {
	X, Y int
}

type Square struct {
	Left, Top, Bottom, Right int
}

func (s *Square) Middle() *Point {
	return &Point{(s.Left + s.Right) / 2, (s.Top + s.Bottom) / 2}

}

type Line struct {
	P1, P2 Point
}

//GetSquaresBisectingLine :
func GetSquaresBisectingLine(s1, s2 Square) *Line {
	s1middle := s1.Middle()
	s2middle := s2.Middle()

	if s1middle == s2middle { //overlap
		return &Line{P1: Point{s1.Left, s1.Top}, P2: Point{s1.Right, s1.Bottom}}
	} //line from one middle to the other
	return &Line{P1: *s1middle, P2: *s2middle}
}
