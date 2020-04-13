package problem0155

// MinStack struct
type MinStack struct {
	nums []int
	mins []int
}

// Constructor func
func Constructor() MinStack {
	return MinStack{}
}

// Push func
func (stack *MinStack) Push(x int) {
	stack.nums = append(stack.nums, x)
	if len(stack.mins) == 0 || x <= stack.GetMin() {
		stack.mins = append(stack.mins, x)
	}
}

// Pop func
func (stack *MinStack) Pop() {
	if stack.Top() == stack.GetMin() {
		stack.mins = stack.mins[:len(stack.mins)-1]
	}
	stack.nums = stack.nums[:len(stack.nums)-1]
}

// Top func
func (stack *MinStack) Top() int {
	return stack.nums[len(stack.nums)-1]
}

// GetMin func
func (stack *MinStack) GetMin() int {
	return stack.mins[len(stack.mins)-1]
}
