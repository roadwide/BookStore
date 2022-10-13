<template>
<!-- 
	div的子元素，也就是class="card"设置了float:left，浮动，会导致父div的高度为零
	好像是子元素浮动到父元素头上了，然后父元素不能自动包裹子元素，导致父div高度为0
	overflow:hidden 可以解决上面的问题
-->
<div style="max-width:720px;overflow:hidden;">
	<div v-show="bookInfo.length" class="card" v-for="book in bookInfo" :key="book.OrderId">
		<img :src="book.pic_url"/>
		<!-- 
		这里文字长度超了会导致布局歪掉，使用text-overflow: ellipsis;使得超出的字符显示为省略号
		/* BOTH of the following are required for text-overflow */
		white-space: nowrap;
		overflow: hidden;
		https://zhuanlan.zhihu.com/p/105511800
		https://codepen.io/catherineliyuankun/pen/gjzLav
		-->
		<p class="card-text">书名:{{book.name}}</p> 
		<p class="card-text">价格:{{book.price}}</p> 
		<p class="card-text">用户:{{book.user_id}}</p> 
	</div>
</div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import axios from 'axios'

const url = 'http://localhost:8081/book/info'
let bookInfo = reactive<any>([])
axios.get(url).then(
function(response) {
	if (response.data.code === 0) {
		for(let oneBookData of response.data.resp) {    // of是内容，in是索引
			//oneBookData长这样 {OrderId: '20221007160940895414', UserId: '123', BookName: 'test4', BookPrice: 1, PicURL: 'http://127.0.0.1:8081/img/20221007160940895414.png'}
			bookInfo.push(oneBookData)
		}
	} else {
		alert("获取所有书目信息失败！")
	}
}).catch(
function(err) {
	alert("未知错误: " + err.message)
})

</script>

<style scoped>
	.card {
		float: left;
		/* 
			每个元素宽100px，border 1px，padding 4px
			每一行放6个就是(100+1*2+4*2) * 6=720
		*/
		width: 110px;
		border: 1px solid #efefef;
		padding: 4px;
		margin-bottom: 2rem;
		text-align: center;
	}

	.card > img {
		margin-bottom: .75rem;
		width: 100px;
		height: 100px;
	}

	.card-text {
		font-size: 85%;
		overflow: hidden;
		white-space: nowrap;
		text-overflow: ellipsis;
	}
</style>