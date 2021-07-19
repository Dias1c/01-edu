package main

import (
	"fmt"
	"strings"
	"student"

	"github.com/01-edu/z01"
)

// import "strconv"

var ColorReset = "\033[0m"
var ColorRed = "\033[31m"
var ColorGreen = "\033[32m"
var ColorBlue = "\033[34m"
var ColorPurple = "\033[35m"
var ColorCyan = "\033[36m"
var ColorWhite = "\033[37m"

func main() {
	TestAtoi()
	// TestRecursivePower()
	// TestPrintCombN()
	// TestPrintNbrBase()
	// TestAtoiBase()
	// TestSplitWhiteSpaces()
	// TestConvertBase()
	// TestSplit()
	// TestActiveBits()
	// TestAdvancedSortArray()
	// TestSortListInsert()
	// TestSortedListMerge()
	// TestListRemoveIf()
	// TestBTreeTransplant()
	// TestBTreeApplyByLevel()
	// TestBTreeDeleteNode()
	// TestAny()
}

func TestAny() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "6")
	fmt.Println("Level: ", student.GetMaxLevelOfTree(root))
}

func TestBTreeDeleteNode() {
	// root := &student.TreeNode{Data: "4"}
	// student.BTreeInsertData(root, "1")
	// student.BTreeInsertData(root, "7")
	// student.BTreeInsertData(root, "5")
	// node := student.BTreeSearchItem(root, "4")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// fmt.Println("By Level")
	// student.BTreeApplyByLevel(root, fmt.Println)
	root := &student.TreeNode{Data: "50"}
	student.BTreeInsertData(root, "20")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "30")
	student.BTreeInsertData(root, "40")

	student.BTreeInsertData(root, "90")
	student.BTreeInsertData(root, "70")
	student.BTreeInsertData(root, "80")
	student.BTreeInsertData(root, "60")

	student.BTreeInsertData(root, "96")
	student.BTreeInsertData(root, "99")
	node := student.BTreeSearchItem(root, "90")
	fmt.Println("Before delete:")
	IsSorted := student.BTreeIsBinary(root)
	fmt.Println("IsSorted: ", IsSorted)
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	IsSorted = student.BTreeIsBinary(root)
	fmt.Println("IsSorted: ", IsSorted)
	student.BTreeApplyInorder(root, fmt.Println)
	fmt.Println("By Level")
	student.BTreeApplyByLevel(root, fmt.Println)
}

//                  5
//           2             9
//        1     3       7     10
//                4   6   8      11
//
// =================================
//                  5
//           2           7
//        1     3      6   8
//                4         10
//                             11
// =================================
func TestBTreeApplyByLevel() {
	root := &student.TreeNode{Data: "4"}

	student.BTreeInsertData(root, "2")
	student.BTreeInsertData(root, "3")
	student.BTreeInsertData(root, "1")

	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "6")
	student.BTreeApplyByLevel(root, fmt.Println)
	fmt.Println("-=-=-=-=-=-=-=-=-=-")
	student.BTreeApplyByLevel(nil, fmt.Println)
	fmt.Println("-=-=-=-=-=-=-=-=-=-")
	root = &student.TreeNode{}
	student.BTreeApplyByLevel(root, fmt.Println)
	fmt.Println("-=-=-=-=-=-=-=-=-=-")
	root = &student.TreeNode{Data: "4"}

	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "2")
	student.BTreeInsertData(root, "3")

	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "6")
	student.BTreeInsertData(root, "7")
	student.BTreeApplyByLevel(root, fmt.Println)
}

// func TestActiveBits() {
// 	fmt.Println(student.ActiveBits(13), 3)
// }

func TestBTreeTransplant() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}

func TestListRemoveIf() {
	link := &student.List{}
	link2 := &student.List{}

	fmt.Println("----normal state----")
	student.ListPushBack(link2, 1)
	PrintList(link2)
	student.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "Hello")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "There")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "How")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "are")
	student.ListPushBack(link, "you")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	PrintList(link)

	student.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
	fmt.Println("Head: ", link.Head)
	fmt.Println("Tail: ", link.Tail)
}

func TestSortedListMerge() {
	var link *student.NodeI
	var link2 *student.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)
	PrintNodes(link)
	PrintNodes(link2)
	PrintNodes(student.SortedListMerge(link2, link))
	//--------------------
	link = nil
	link = listPushBack(link, 1)
	link = listPushBack(link, 10)
	link = listPushBack(link, 11)
	link = listPushBack(link, 20)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)
	link = listPushBack(link, 2)
	PrintNodes(link)

	link = student.SortedListMerge(link, link)
	PrintNodes(link)
}

func TestSortListInsert() {
	var link *student.NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)
	PrintNodes(link)
	link = student.SortListInsert(link, -2)
	PrintNodes(link)
	link = student.SortListInsert(link, 2)
	PrintNodes(link)
}

func TestAdvancedSortArray() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result, student.Compare)

	fmt.Println(result)
}

func TestConvertBase() {
	PrintTitle("TestConvertBase")
	fmt.Println("=-=-=-= MainTest =-=-=-=-=")
	ExceptedAndGetsString("hccocimc", student.ConvertBase("4506C", "0123456789ABCDEF", "choumi"))
	ExceptedAndGetsString("9B611", student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF"))
	ExceptedAndGetsString("111110101101010010", student.ConvertBase("256850", "0123456789", "01"))
	ExceptedAndGetsString("eeone0n", student.ConvertBase("uuhoumo", "choumi", "Zone01"))
	ExceptedAndGetsString("683241", student.ConvertBase("683241", "0123456789", "0123456789"))
	fmt.Println("=-=-=-= ---- =-=-=-=-=")
	ExceptedAndGetsString("43", student.ConvertBase("101011", "01", "0123456789"))
}

func TestSplit() {
	PrintTitle("TestSplit")
	fmt.Println("=-=-=-= MainTest =-=-=-=-=")
	excepteds := [2][]string{
		{
			" Hello how are you? ",
			"The earliest foundations of what would become computer science predate the invention of the modern digital computer",
			"|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,",
			"AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,",
		},
		{
			" ",
			" ",
			"|=choumi=|",
			"AFJ",
		},
	}
	isPassed := true
	for i := range excepteds {
		if !ExceptedAndGetsSliceString(strings.Split(excepteds[0][i], excepteds[1][i]), student.Split(excepteds[0][i], excepteds[1][i])) {
			isPassed = false
		}
	}
	PrintPassedFailed("MainTest:", isPassed)
	fmt.Println("=-=-=-= ---- =-=-=-=-=")
}

func TestSplitWhiteSpaces() {
	PrintTitle("TestSplitWhiteSpaces")
	fmt.Println("=-=-=-= MainTest =-=-=-=-=")
	excepteds := [2][][]string{
		{
			{"Hello", "how", "are", "you?"},
			{"The", "earliest", "foundations", "of", "what", "would", "become", "computer", "science", "predate", "the", "invention", "of", "the", "modern", "digital", "computer"},
		},
		{
			student.SplitWhiteSpaces("Hello how are you?"),
			student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"),
		},
	}
	isPassed := true
	for i := range excepteds {
		if !ExceptedAndGetsSliceString(excepteds[0][i], excepteds[1][i]) {
			isPassed = false
		}
	}
	PrintPassedFailed("MainTest:", isPassed)
	fmt.Println("=-=-=-= ---- =-=-=-=-=")
}

func TestAtoiBase() {
	PrintTitle("TestAtoiBase")
	fmt.Println("=-=-=-= MainTest =-=-=-=-=")
	excepteds := [2][]int{
		{
			12016,
			1,
			0,
			11557277891,
			406772,
			0,
			125,
			125,
			0,
		},
		{
			student.AtoiBase("bcbbbbaab", "abc"),
			student.AtoiBase("0001", "01"),
			student.AtoiBase("00", "01"),
			student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"),
			student.AtoiBase("AAho?Ao", "WhoAmI?"),
			student.AtoiBase("thisinputshouldnotmatter", "abca"),
			student.AtoiBase("125", "0123456789"),
			student.AtoiBase("uoi", "choumi"),
			student.AtoiBase("bbbbbab", "-ab"),
		},
	}
	isPassed := true
	for i := range excepteds {
		if !ExceptedAndGetsInt(excepteds[0][i], excepteds[1][i]) {
			isPassed = false
		}
	}
	PrintPassedFailed("MainTest:", isPassed)
	fmt.Println("=-=-=-= ---- =-=-=-=-=")
	ExceptedAndGetsInt(125, student.AtoiBase("125", "0123456789"))
	ExceptedAndGetsInt(125, student.AtoiBase("1111101", "01"))
	ExceptedAndGetsInt(125, student.AtoiBase("7D", "0123456789ABCDEF"))
	ExceptedAndGetsInt(125, student.AtoiBase("uoi", "choumi"))
	ExceptedAndGetsInt(0, student.AtoiBase("bbbbbab", "-ab"))
}

func TestPrintNbrBase() {
	PrintTitle("TestPrintNbrBase")
	fmt.Println("=-=-=-= MainTest =-=-=-=-=")
	student.PrintNbrBase(919617, "01")
	fmt.Print(" == 11100000100001000001")
	z01.PrintRune('\n')
	student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	fmt.Print(" == HIDAHI")
	z01.PrintRune('\n')
	student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	fmt.Print(" == -MssiD")
	z01.PrintRune('\n')
	student.PrintNbrBase(-661165, "1")
	fmt.Print(" == NV")
	z01.PrintRune('\n')
	student.PrintNbrBase(-861737, "Zone01Zone01")
	fmt.Print(" == NV")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	fmt.Print(" == 7D")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	fmt.Print(" == -uoi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "-ab")
	fmt.Print(" == NV")
	z01.PrintRune('\n')
	student.PrintNbrBase(-9223372036854775808, "0123456789")
	fmt.Print(" == -9223372036854775808")
	z01.PrintRune('\n')
	fmt.Println("=-=-=-= MyTests =-=-=-=-=")
	student.PrintNbrBase(125, "0123456789")
	fmt.Print(" == 125")
	z01.PrintRune('\n')
	student.PrintNbrBase(0, "0123456789")
	fmt.Print(" == 0")
	z01.PrintRune('\n')
	student.PrintNbrBase(0, "chaomi")
	fmt.Print(" == c")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "01")
	fmt.Print(" == -1111101")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "aa")
	fmt.Print(" == NV")
	z01.PrintRune('\n')
	student.PrintNbrBase(9223372036854775807, "0123456789")
	fmt.Print(" == 9223372036854775807")
	z01.PrintRune('\n')
	student.PrintNbrBase(-9223372036854775808, "0123456789")
	fmt.Print(" == -9223372036854775808")
	z01.PrintRune('\n')
}

func TestPrintCombN() {
	PrintTitle("TestPrintCombN")
	student.PrintCombN(1)
	fmt.Println("\n--------")
	student.PrintCombN(2)
	fmt.Println("\n--------")
	student.PrintCombN(3)
	fmt.Println("\n--------")
	student.PrintCombN(4)
	fmt.Println("\n--------")
	student.PrintCombN(5)
	fmt.Println("\n--------")
	student.PrintCombN(6)
	fmt.Println("\n--------")
	student.PrintCombN(7)
	fmt.Println("\n--------")
	student.PrintCombN(8)
	fmt.Println("\n--------")
	student.PrintCombN(9)
	fmt.Println("\n-------- Unreal Tests (Should Be Empty)")
	student.PrintCombN(0)
	fmt.Println("\n--------")
	student.PrintCombN(10)
	fmt.Println("\n--------")
}

func TestRecursivePower() {
	PrintTitle("TestRecursivePower")
	fmt.Println("=-=-=-= MainTest =-=-=-=-=")
	ExceptedAndGetsInt(0, student.RecursivePower(-7, -2))
	ExceptedAndGetsInt(0, student.RecursivePower(-8, -7))
	ExceptedAndGetsInt(65536, student.RecursivePower(4, 8))
	ExceptedAndGetsInt(1, student.RecursivePower(1, 3))
	ExceptedAndGetsInt(-1, student.RecursivePower(-1, 1))
	ExceptedAndGetsInt(-7776, student.RecursivePower(-6, 5))
	fmt.Println("=-=-=-= MyTests =-=-=-=-=")
	ExceptedAndGetsInt(64, student.RecursivePower(4, 3))
	ExceptedAndGetsInt(9, student.RecursivePower(9, 1))
	ExceptedAndGetsInt(1, student.RecursivePower(9, 0))
	ExceptedAndGetsInt(0, student.RecursivePower(9, -1))
	ExceptedAndGetsInt(0, student.RecursivePower(9, -2147483648))
	ExceptedAndGetsInt(0, student.RecursivePower(9, -9223372036854775808))
}

func TestAtoi() {
	PrintTitle("TestAtoi")
	fmt.Println("=-=-=-= MainTest =-=-=-=-=")
	ExceptedAndGetsInt(0, student.Atoi(""))
	ExceptedAndGetsInt(0, student.Atoi("-"))
	ExceptedAndGetsInt(0, student.Atoi("--123"))
	ExceptedAndGetsInt(1, student.Atoi("1"))
	ExceptedAndGetsInt(-3, student.Atoi("-3"))
	ExceptedAndGetsInt(8292, student.Atoi("8292"))
	ExceptedAndGetsInt(9223372036854775807, student.Atoi("9223372036854775807"))
	ExceptedAndGetsInt(-9223372036854775808, student.Atoi("-9223372036854775808"))
	fmt.Println("=-=-=-= MyTests =-=-=-=-=")
	ExceptedAndGetsInt(12345, student.Atoi("12345"))
	ExceptedAndGetsInt(12345, student.Atoi("0000000012345")) // This
	ExceptedAndGetsInt(0, student.Atoi("012 345"))
	ExceptedAndGetsInt(0, student.Atoi("Hello World!"))
	ExceptedAndGetsInt(0, student.Atoi(" 1"))
	ExceptedAndGetsInt(1234, student.Atoi("+1234"))
	ExceptedAndGetsInt(-1234, student.Atoi("-1234"))
	ExceptedAndGetsInt(0, student.Atoi("++1234"))
	ExceptedAndGetsInt(0, student.Atoi("--1234"))
	ExceptedAndGetsInt(-2147483648, student.Atoi("-2147483648"))
	ExceptedAndGetsInt(2147483647, student.Atoi("2147483647"))
	ExceptedAndGetsInt(2147483648, student.Atoi("2147483648"))
	ExceptedAndGetsInt(-9223372036854775808, student.Atoi("-9223372036854775808"))
	ExceptedAndGetsInt(9223372036854775807, student.Atoi("9223372036854775807"))
	ExceptedAndGetsInt(9223372036854775807, student.Atoi("9223372036854775808"))
	ExceptedAndGetsInt(-9223372036854775808, student.Atoi("-9223372036854775809"))
	ExceptedAndGetsInt(-9223372036854775808, student.Atoi("-92233720368547758000000009"))
	ExceptedAndGetsInt(9223372036854775807, student.Atoi("0000000000009223372036854775807")) // This
	ExceptedAndGetsInt(-9223372036854775807, student.Atoi("-0000000000009223372036854775807"))
	ExceptedAndGetsInt(9223372036854775807, student.Atoi("13456563454455454535440"))
	ExceptedAndGetsInt(0, student.Atoi("0-5"))
	ExceptedAndGetsInt(9223372036854775807, student.Atoi("90000000000000000000"))
	ExceptedAndGetsInt(9223372036854775807, student.Atoi("50000000000000000000"))
}

func ExceptedAndGetsSliceString(excepted, res []string) bool {
	if len(excepted) != len(res) {
		fmt.Print(string(ColorRed))
		fmt.Println("Lens is not Contains")
	}
	isCompare := true
	for i := range excepted {
		if excepted[i] != res[i] {
			isCompare = false
			fmt.Print(string(ColorRed))
			fmt.Printf("excepted[%v] != res[%v]\n", i, i)
			fmt.Printf("'%v' != '%v'\n", excepted[i], res[i])
		}
	}
	fmt.Print(string(ColorReset))
	if !isCompare {
		fmt.Println(excepted)
		fmt.Println(res)
		fmt.Println()
	}
	return isCompare
}

func ExceptedAndGetsInt(excepted, res int) bool {
	color, equal, isCompare := ColorRed, "!=", excepted == res
	if isCompare {
		color, equal = ColorGreen, "=="
	}
	if !isCompare {
		fmt.Print("Res: ")
		fmt.Print(string(color))
		fmt.Printf("%v\t", isCompare)
		fmt.Print(string(ColorReset))
		fmt.Printf("Exepted: %v %v %v\n", excepted, equal, res)
	}
	return isCompare
}

func ExceptedAndGetsString(excepted, res string) bool {
	color, equal, isCompare := ColorRed, "!=", excepted == res
	if isCompare {
		color, equal = ColorGreen, "=="
	}
	if !isCompare {
		fmt.Print("Res: ")
		fmt.Print(string(color))
		fmt.Printf("%v\t", isCompare)
		fmt.Print(string(ColorReset))
		fmt.Printf("Exepted: %v %v %v\n", excepted, equal, res)
	}
	return isCompare
}

func PrintTitle(title string) {
	fmt.Print(string(ColorCyan))
	fmt.Println(title)
	fmt.Print(string(ColorReset))
}

func PrintPassedFailed(title string, isPassed bool) {
	fmt.Print(title)
	if isPassed {
		fmt.Print(string(ColorGreen))
		fmt.Println("PASSED")
	} else {
		fmt.Print(string(ColorRed))
		fmt.Println("FAILED")
	}
	fmt.Print(string(ColorReset))
}

func PrintNodes(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func PrintList(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}
