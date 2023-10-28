package main

import (
  "fmt"
)

type tree struct {
  value int
  left *tree
  right *tree
}

func addNode(root *tree, value int) {
  for {
    if value < root.value {
      if root.left == nil {
        root.left = &tree{value, nil, nil}
        return
      } else {
        root = root.left
      }
    } else if value > root.value {
      if root.right == nil {
        root.right = &tree{value, nil, nil}
      } else {
        root = root.right
      }
    // Duplicate value. Don't add
    } else {
      return 
    }
  }
}

func addNodesFromSlice(root *tree, values []int) {
  for _, v := range values {
    addNode(root, v) 
  }
}

func inOrder(root *tree) {
  if root.left != nil {
    inOrder(root.left)
  }

  fmt.Println(root.value)

  if root.right != nil {
    inOrder(root.right)
  }
}

func main() {
  root := &tree{8, nil, nil}
  fmt.Println(root)

  addNode(root, 10)
  addNode(root, 3)
  addNode(root, 13)
  addNode(root, 7)

  fmt.Println("Tree one")
  inOrder(root)
  
  root2 := &tree{5, nil, nil}

  nodesToAdd := []int{10, -2, 17, 22, 0, 5, 100, -11, -313}
  addNodesFromSlice(root2, nodesToAdd)

  fmt.Println("\nTree two")
  inOrder(root2)
}
