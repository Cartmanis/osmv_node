package osmv_node

type Osmv struct {
	IshV int64
	Npar string
	Dat string
}

type Node struct {
	Parent *Node
	Childerns []*Node
	Sibling []*Node
	IshV int64
	Npar string
	Dat string
}

func GetNodesFromOsmv(listOsmv []*Osmv) []*Node {
	nodeList := make([]*Node, 0)

	for _, v := range listOsmv {
		node := &Node{nil, nil, nil, v.IshV, v.Npar, v.Dat}
		nodeList = append(nodeList, node)
	}
	return getOsmv(nodeList)
}

func getOsmv(listNode []*Node) []*Node {
	maxIshv := getMaxIshv(listNode)
	if maxIshv == 0 {
		fillSibling(listNode)
		return listNode
	}

	predIndex := 0
	siblingList := make([]*Node, 0)
	for i, v := range listNode {
		if v.IshV == maxIshv {
			listNode[predIndex].Childerns = append(listNode[predIndex].Childerns, &Node{ listNode[predIndex], v.Childerns, nil,v.IshV, v.Npar, v.Dat})
			siblingList = append(siblingList, v)
			continue
		}
		fillSibling(listNode[predIndex].Childerns)
		siblingList = make([]*Node, 0)
		predIndex = i
	}
	//удаляем элементы, которые уже являются childern другого элемента, чтобы далее вызвать рекурсивно более маленькую сущность
	return getOsmv(filterListNode(listNode, maxIshv))
}

func getMaxIshv(listNode []*Node) int64 {
	var maxIshv int64
	for _, v := range listNode {
		if v.IshV > maxIshv {
			maxIshv = v.IshV
		}
	}
	return maxIshv
}

func filterListNode(listNode []*Node, max int64) []*Node {
	filterList := make([]*Node, 0)
	for _, v := range listNode {
		if v.IshV != max {
			filterList = append(filterList, v)
		}
	}
	return filterList
}

func fillSibling(nodeList []*Node) {
	for _, n := range nodeList {
		filterList := make([]*Node,0)
		for _, nn := range  nodeList {
			if nn.Npar == n.Npar && nn.Dat == n.Dat && nn.IshV == n.IshV {
				continue
			}
			filterList = append(filterList, nn)
		}
		n.Sibling = filterList
	}
}