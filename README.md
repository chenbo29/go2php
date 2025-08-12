# go2php

[![Go](https://github.com/chenbo29/go2php/actions/workflows/go.yml/badge.svg)](https://github.com/chenbo29/go2php/actions/workflows/go.yml)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/chenbo29/go2php)
[![codecov](https://codecov.io/gh/chenbo29/go2php/branch/main/graph/badge.svg?token=7P4UHIRF1K)](https://codecov.io/gh/chenbo29/go2php)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenbo29/go2php)](https://goreportcard.com/report/github.com/chenbo29/go2php)

GoLang alternatives for PHP functions

## Array Functions
## https://www.php.net/manual/en/ref.array.php
## 进度
* [✔] array_change_key_case — Changes the case of all keys in an array — 将数组中的所有键名修改为全大写或小写
* [x] array_chunk — Split an array into chunks - 将一个数组分割成多个
* [x] array_column — Return the values from a single column in the input array - 返回输入数组中指定列的值
* [x] array_combine — Creates an array by using one array for keys and another for its values- 创建一个数组，用一个数组的值作为其键名，另一个数组的值作为其值
* [x] array_count_values — Counts all the values of an array - 统计数组中所有的值
* [x] array_diff_assoc — Computes the difference of arrays with additional index check - 带索引检查计算数组的差集
* [ ] array_diff_key — 使用键名比较计算数组的差集
* [ ] array_diff_uassoc — 用用户提供的回调函数做索引检查来计算数组的差集
* [ ] array_diff_ukey — 用回调函数对键名比较计算数组的差集
* [x] array_diff — Computes the difference of arrays - 计算数组的差集
* [ ] array_fill_keys — 使用指定的键和值填充数组
* [ ] array_fill — 用给定的值填充数组
* [ ] array_filter — 使用回调函数过滤数组的元素
* [ ] array_flip — 交换数组中的键和值
* [ ] array_intersect_assoc — 带索引检查计算数组的交集
* [ ] array_intersect_key — 使用键名比较计算数组的交集
* [ ] array_intersect_uassoc — 带索引检查计算数组的交集，用回调函数比较索引
* [ ] array_intersect_ukey — 在键名上使用回调函数来比较计算数组的交集
* [ ] array_intersect — 计算数组的交集
* [ ] ❌array_is_list — 判断给定的 array 是否为 list
* [ ] array_key_exists — 检查数组里是否有指定的键名或索引
* [ ] array_key_first — 获取指定数组的第一个键
* [ ] array_key_last — 获取一个数组的最后一个键值
* [x] array_keys — 返回数组中部分的或所有的键名
* [ ] array_map — 为数组的每个元素应用回调函数
* [ ] array_merge_recursive — 递归地合并一个或多个数组
* [x] array_merge — 合并一个或多个数组
* [ ] array_multisort — 对多个数组或多维数组进行排序
* [ ] array_pad — 以指定长度将一个值填充进数组
* [ ] array_pop — 弹出数组最后一个单元（出栈）
* [ ] array_product — 计算数组中所有值的乘积
* [x] array_push — 将一个或多个单元压入数组的末尾（入栈）
* [ ] array_rand — 从数组中随机取出一个或多个随机键
* [ ] array_reduce — 用回调函数迭代地将数组简化为单一的值
* [ ] array_replace_recursive — 使用传递的数组递归替换第一个数组的元素
* [ ] array_replace — 使用传递的数组替换第一个数组的元素
* [ ] array_reverse — 返回单元顺序相反的数组
* [ ] array_search — 在数组中搜索给定的值，如果成功则返回首个相应的键名
* [ ] array_shift — 将数组开头的单元移出数组
* [ ] array_slice — 从数组中取出一段
* [ ] array_splice — 去掉数组中的某一部分并用其它值取代
* [ ] array_sum — 对数组中所有值求和
* [ ] array_udiff_assoc — 带索引检查计算数组的差集，用回调函数比较数据
* [ ] array_udiff_uassoc — 带索引检查计算数组的差集，用回调函数比较数据和索引
* [ ] array_udiff — 用回调函数比较数据来计算数组的差集
* [ ] array_uintersect_assoc — 带索引检查计算数组的交集，用回调函数比较数据
* [ ] array_uintersect_uassoc — 带索引检查计算数组的交集，用单独的回调函数比较数据和索引
* [ ] array_uintersect — 计算数组的交集，用回调函数比较数据
* [ ] array_unique — 移除数组中重复的值
* [ ] array_unshift — 在数组开头插入一个或多个单元
* [ ] array_values — 返回数组中所有的值
* [ ] array_walk_recursive — 对数组中的每个成员递归地应用用户函数
* [ ] array_walk — 使用用户自定义函数对数组中的每个元素做回调处理
* [ ] array — 新建一个数组
* [ ] arsort — 对数组进行降向排序并保持索引关系
* [ ] asort — 对数组进行升序排序并保持索引关系
* [ ] compact — 建立一个数组，包括变量名和它们的值
* [ ] count — 统计数组、Countable 对象中所有元素的数量
* [ ] current — 返回数组中的当前值
* [ ] each — 返回数组中当前的键／值对并将数组指针向前移动一步
* [ ] end — 将数组的内部指针指向最后一个单元
* [ ] extract — 从数组中将变量导入到当前的符号表
* [x] in_array — 检查数组中是否存在某个值
* [ ] key_exists — 别名 array_key_exists
* [ ] key — 从关联数组中取得键名
* [ ] krsort — 对数组按照键名逆向排序
* [ ] ksort — 对数组根据键名升序排序
* [ ] list — 把数组中的值赋给一组变量
* [ ] natcasesort — 用“自然排序”算法对数组进行不区分大小写字母的排序
* [ ] natsort — 用“自然排序”算法对数组排序
* [ ] next — 将数组中的内部指针向前移动一位
* [ ] pos — current 的别名
* [ ] prev — 将数组的内部指针倒回一位
* [ ] range — 根据范围创建数组，包含指定的元素
* [ ] reset — 将数组的内部指针指向第一个单元
* [ ] rsort — 对数组降序排序
* [ ] shuffle — 打乱数组
* [ ] sizeof — count 的别名
* [ ] sort — 对数组升序排序
* [ ] uasort — 使用用户自定义的比较函数，保持索引和值的对应关系，原地排序 array。
* [ ] uksort — 使用用户自定义的比较函数对数组中的键名进行排序
* [ ] usort — 使用用户自定义的比较函数对数组中的值进行排序