package sorting

func Revert(arr []int) []int {
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

type binaryTree struct {
	value       int
	left, right *binaryTree
}

func add(tree *binaryTree, value int) *binaryTree {
	if tree == nil {
		tree = new(binaryTree)
		tree.value = value
		return tree
	}

	if value <= tree.value {
		tree.left = add(tree.left, value)
		return tree
	}

	tree.right = add(tree.right, value)
	return tree
}

func appendSlice(values []int, tree *binaryTree) []int {
	if tree == nil {
		return values
	}
	values = appendSlice(values, tree.left)
	values = append(values, tree.value)
	values = appendSlice(values, tree.right)
	return values
}

func InsertionSort_BinaryTree(originSlice []int) (sortedSlice []int) {
	var binTree *binaryTree
	for _, value := range originSlice {
		binTree = add(binTree, value)
	}

	sortedSlice = []int{}
	sortedSlice = appendSlice(sortedSlice, binTree)
	return sortedSlice
}
