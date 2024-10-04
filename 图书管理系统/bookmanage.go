package main

import (
	"fmt"
)

type Book struct {
	ID     int
	Name   string
	Author string
}

type manager struct {
	idnumber int
	bor      [][]string //也可用结构体切片，操作与booklist相似
	booklist []Book
}

// 添加
func (m *manager) add(book *Book) {
	fmt.Println("请输入要添加的书籍数目")
	addnum := 0
	fmt.Scanln(&addnum)
	if addnum > 0 {
		for i := 0; i < addnum; i++ {
			fmt.Println("请输入书名：")
			fmt.Scanln(&book.Name)
			fmt.Println("请输入作者：")
			fmt.Scanln(&book.Author)
			book.ID = m.idnumber
			m.booklist = append(m.booklist, *book)
			m.idnumber++
		}
	} else {
		fmt.Println("输入有误")
	}
}

// 删除
func (m *manager) delete() {
	fmt.Println("请输入要删除的书籍数目")
	delnum := 0
	fmt.Scanln(&delnum)
	if delnum <= 0 || delnum > len(m.booklist) {
		fmt.Println("输入有误")
		return
	}
	for i := 0; i < delnum; i++ {

		fmt.Println("请输入要删除的书籍的名称:")
		var name string
		fmt.Scanln(&name)
		var del bool = false
		for j := 0; j < len(m.booklist); j++ {
			if m.booklist[j].Name == name {
				del = true
				for k := j; k < len(m.booklist)-1; k++ {
					m.booklist[k] = m.booklist[k+1]
					//不用append（booklist[:k],booklist[k+1:]）--会导致数组下标越界
				}
				fmt.Println("删除成功！")
				break
			}
		}
		m.booklist = m.booklist[:len(m.booklist)-1] //删除最后一个元素
		if !del {
			fmt.Println("未找到该书籍，请重新输入")
		}

	}
}

// 显示所有图书
func (m *manager) show() {
	fmt.Println("书籍列表：")
	for i := 0; i < len(m.booklist); i++ {
		fmt.Println("书名：", m.booklist[i].Name, " 作者：", m.booklist[i].Author, "ID:", m.booklist[i].ID)
	}
}

// 借阅
func (m *manager) borrow() {

	fmt.Println("每人总共最多可借三本书")
	fmt.Println("请输入借书人的姓名：")
	name := ""
	fmt.Scanln(&name)
	length := 0
	limit := 3
	position := 0
	fmt.Println(m.bor)
	fmt.Println(len(m.bor))
	//判断是否有借阅记录
	if len(m.bor) > 0 {
		for k := 0; k < len(m.bor); k++ {
			if m.bor[k][0] == name {
				length = len(m.bor[k])
				limit = 3 - (length - 1) //length-1表示已借阅的书籍数量
				position = k             //定位该借阅人信息在bor切片中的位置
				break
			}
		}
	}
	if limit == 0 {
		fmt.Println("该人已借阅三本书，无法再借阅,请先归还书籍")
		return
	} else {
		fmt.Println("您本次可以借阅的书籍数量最多为：", limit, "本")
	start:
		bornum := 0
		fmt.Println("请输入借书数量：")
		fmt.Scanln(&bornum)
		if bornum > limit {
			fmt.Println("借书数量超过可借数量，请重新输入")
			goto start
		} else {
			if length == 0 {
				addbor := make([]string, 1)
				for i := 0; i < bornum; i++ {
				start2:
					fmt.Println("请输入借阅的书籍名称:")
					book := ""
					fmt.Scanln(&book)
					var exist bool = false
					for j := 0; j < len(m.booklist); j++ {
						if m.booklist[j].Name == book {
							exist = true
							break
						}
					}
					if !exist {
						fmt.Println("该书籍不存在，请重新输入")
						i--
						continue //等效于goto start2
					}
					for _, value1 := range m.bor {
						for _, value2 := range value1 {
							if value2 == book {
								fmt.Println("该书籍已被借阅，请重新输入")

								goto start2
							}
						}
					}
					addbor[0] = name
					addbor = append(addbor, book)
					fmt.Println("借阅成功！")
				}
				m.bor = append(m.bor, addbor) //放在循环外，避免添加了多个切片
			} else {
				for i := 0; i < bornum; i++ {
				start3:
					fmt.Println("请输入借阅的书籍名称:")
					book := ""
					fmt.Scanln(&book)
					var exist bool = false
					for j := 0; j < len(m.booklist); j++ {
						if m.booklist[j].Name == book {
							exist = true
							break
						}
					}
					if !exist {
						fmt.Println("该书籍不存在，请重新输入")
						i--
						continue
					}
					for _, value1 := range m.bor {
						for _, value2 := range value1 {
							if value2 == book {
								fmt.Println("该书籍已被借阅，请重新输入")

								goto start3
							}
						}
					}
					m.bor[position] = append(m.bor[position], book)
					fmt.Println("借阅成功！")
				}
			}
		}

	}

}

// 查看借阅信息
func (m *manager) check() {
	fmt.Println("请输入查询人的名称：")
	checkname := ""
	fmt.Scanln(&checkname)
	fmt.Println(len(m.bor))
	for i := 0; i < len(m.bor); i++ {
		if m.bor[i][0] == checkname {
			fmt.Println(checkname, "的借阅情况的书籍如下：")
			for j := 1; j < len(m.bor[i]); j++ {
				fmt.Println("书名：", m.bor[i][j])
			}
			return
		}

	}
	fmt.Println("未借阅任何书籍")
}

// 归还
func (m *manager) returnback() {
	fmt.Println("请输入还书人的姓名：")
	name := ""
	fmt.Scanln(&name)
	isexist := false
	for i := 0; i < len(m.bor); i++ {
		if m.bor[i][0] == name {
			isexist = true
		}
		if isexist {
		start:
			returnnum := 0
			fmt.Println("请输入还书的数量：")
			fmt.Scanln(&returnnum)
			if returnnum > len(m.bor[i])-1 {
				fmt.Println("还书数量超过借阅数量，请重新输入")
				goto start
			}
			for j := 0; j < returnnum; j++ {
				fmt.Println("请输入归还书籍的名称：")
				bookname := ""
				fmt.Scanln(&bookname)
				templist := make([]string, 0) //创建一个临时切片，用于存储还书后剩余的书籍
				for _, value := range m.bor[i] {
					// bor[i]=append(bor[i][:j],bor[i][j+1:]...)--此方法会出现下标越界的问题

					if value != bookname {
						templist = append(templist, value)

					}
				}
				if len(templist) == len(m.bor[i]) {
					fmt.Println("未借阅该书籍")
				} else {

					m.bor[i] = templist
					fmt.Println("归还成功！")
				}

			}
		} else {
			fmt.Println("无借书记录")
		}
	}

}

// 菜单
func menu() {
	fmt.Println("-----欢迎使用图书管理系统-----")
	fmt.Println("1.添加书籍")
	fmt.Println("2.删除书籍")
	fmt.Println("3.显示所有书籍")
	fmt.Println("4.借阅书籍")
	fmt.Println("5.查看借阅情况")
	fmt.Println("6.归还书籍")
	fmt.Println("7.退出系统")
	fmt.Println("----------------------------")

}
func main() {
	var book *Book = &Book{}
	m := manager{}
	for {
		menu()
		fmt.Println("请输入选项：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			m.add(book)
		case 2:
			m.delete()
		case 3:
			m.show()
		case 4:
			m.borrow()
		case 5:
			m.check()
		case 6:
			m.returnback()
		case 7:
			fmt.Println("欢迎下次使用！")
			return
		default:
			fmt.Println("输入错误，请重新输入")
		}
	}
}
