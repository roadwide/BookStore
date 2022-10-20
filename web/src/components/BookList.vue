<template>
<div style="max-width:720px;overflow:hidden;">
	<div v-show="props.bookData.length" class="card" v-for="book in props.bookData" :key="book.id">
		<img :src="book.pic_url"/>
		<p class="card-text">书名:{{book.name}}</p>  
		<p class="card-text">价格:{{book.price}}</p> 
		<p class="card-text">用户:{{book.user_id}}</p> 
		<button v-if="props.isMyBookList" @click="deleteBook(book.id, userInfo.token)">删除</button>
	</div>
</div>
</template>

<script setup lang="ts">
import {defineProps} from 'vue'
import axios from 'axios'
const props = defineProps(['bookData', 'isMyBookList'])
const userInfo = JSON.parse(localStorage.getItem("userInfo")!)
function deleteBook(bookID:number, token:string) {
	const formData = {
		book_id: bookID,
		token: token
	}
	axios.post(process.env.VUE_APP_BASE_API+'/book/delete', formData).then(
        function(response) {
          if (response.data.code === 0) {
            alert("删除成功！")
			// 页面直接刷新，闪烁，用户体验不好。后期再改
			window.location.reload();
          } else {
            alert("删除失败！")
          }
        }
      ).catch(
        function(err) {
          alert("未知错误: " + err.message)
        }
      )
}
</script>

<style scoped>
	.card {
		float: left;
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