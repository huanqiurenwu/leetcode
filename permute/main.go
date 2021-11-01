package main

import "fmt"

func main(){
	fmt.Println(len(permute([]int{1,2,3})))
	input := []int{1,2,3,4,5}
	remove(&input, 5)
	input = append(input, 6,7,8)
	fmt.Println(cap(input), " ", len(input))
	fmt.Println(input[10:10])
}


func permute(input []int) [][]int{
	trackSet := [][]int{}
	var backtrack func([]int, []int)
	backtrack = func (track []int, selective []int){
		if len(track) == len(selective){
			trackSet = append(trackSet, track)
			return
		}
		for _, option := range selective{
			if contains(track, option) {
				continue
			}
			track = append(track, option)
			backtrack(track, selective)
			track = track[:len(track) - 1]
		}
	}
	backtrack([]int{}, input)
	return trackSet
}


func contains(list []int, value int) bool{
	for _,v := range list{
		if v == value{
			return true
		}
	}
	return false
}

func remove(input *[]int, value int){
	for i,v := range *input{
		if v == value {
			*input = append((*input)[:i], (*input)[i+1:]...)
		}
	}
}

func returnCheck(in int)(int, b int){
	int = 99
	return int, b
}