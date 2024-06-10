class Stack:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)
    
    def pop(self):
        if not self.is_empty():
            return self.items.pop()

    def is_empty(self):
        return self.items == []

    def peek(self):
        if not self.is_empty():
            return self.items[-1]

class Solution:
    def isValid(self, s: str) -> bool:
        stack = Stack()
        for v in s:
            if v == '(' or v == '[' or v == '{':
                stack.push(v)
            elif v == ')':
                if stack.peek() == '(':
                    stack.pop()
                else:
                    return False
            elif v == ']':
                if stack.peek() == '[':
                    stack.pop()
                else:
                    return False
            elif v == '}':
                if stack.peek() == '{':
                    stack.pop()
                else:
                    return False
        if stack.is_empty():
            return True
        return False