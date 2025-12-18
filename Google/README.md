# Problem1

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7,8, 9, 14] and k of 17, return true since 10 + 7 is 17.

Note: do this in one pass.

# Problem 2

Given the root to a binary tree, implement *serialize(root)*, which serializes the tree into a string, and *deserialize(s)*, which deserializes the string back into the tree.

For example, given the following Node class

*class Node:*
    *def __init__(self, val, left=None, right=None):*
        *self.val = val*
        *self.left = left*
        *self.right = right*

The following test should pass:

*node = Node('root', Node('left', Node('left.left')), Node('right'))*
*assert deserialize(serialize(node)).left.left.val == 'left.left'*