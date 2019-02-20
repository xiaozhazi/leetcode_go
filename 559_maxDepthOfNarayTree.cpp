/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
public:
    int maxDepth(Node* root) {
        int depth = 0;
        if (root == NULL) {
            return depth;
        }

        for(auto n : root->children)
            depth = max(depth, maxDepth(n));
        return depth + 1;
    }
};