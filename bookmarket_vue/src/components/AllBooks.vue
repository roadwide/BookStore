<template>
  <h1>All Books</h1>
  <div class="row">
        <!-- 展示用户列表 -->
        <div v-show="bookInfo.length" class="card" v-for="book in bookInfo" :key="book.OrderId">
			<img :src="book.PicURL" style="width: 100px" />
            <p class="card-text">{{book.BookName}}</p> 
        </div>
	</div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import axios from 'axios'

const url = 'http://localhost:8081/getAllBookInfo'
let bookInfo = reactive<any>([])
axios.get(url).then(
function(response) {
	if (response.data.message === "ok") {
		for(let oneBookData of response.data.data) {    // of是内容，in是索引
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
		width: 33.333%;
		padding: .75rem;
		margin-bottom: 2rem;
		border: 1px solid #efefef;
		text-align: center;
	}

	.card > img {
		margin-bottom: .75rem;
		border-radius: 100px;
	}

	.card-text {
		font-size: 85%;
	}
</style>