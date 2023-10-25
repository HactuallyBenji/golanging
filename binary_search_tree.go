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

  inOrder(root)
}
