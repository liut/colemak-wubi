

转换 `qwerty` 布局的五笔码表到 `colemak-dh` 布局。

美中不足的是，由于colemak分号位置的变更，暂只能将原P键位改为分号位（字母O键位）

注意：只在可以改键位的键盘下测试过。


先准备好五笔的码表文件，以清歌为例：

* wb_98_table.txt：带词语词库
* wb_98_reverse_table.txt：反查词库

```bash
go run wbcm.go -i wb_98_table.txt -o cwb_98_table.txt
go run wbcm.go -i wb_98_reverse_table.txt -o cwb_98_reverse_table.txt
```

转换完成后，在清歌中选择并生成索引即可使用
