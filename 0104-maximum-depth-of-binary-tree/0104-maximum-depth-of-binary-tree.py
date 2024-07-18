# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

    max: int = 1

    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        
        self.recursive(1, root)

        return self.max

    def recursive(self, depth: int, node: Optional[TreeNode]):
        if node and node.left:
            self.recursive(depth + 1, node.left)
            self.max = max(self.max, depth + 1)
        if node and node.right:
            self.recursive(depth + 1, node.right)
            self.max = max(self.max, depth + 1)