<template>
<div style="max-width:720px;overflow:hidden;">
	<div v-show="props.bookData.length" class="card" v-for="book in props.bookData" :key="book.id">
		<img :src="book.pic_url"/>
		<p class="card-text">书名:{{book.name}}</p> 
		<p class="card-text">价格:{{book.price}}</p> 
		<p class="card-text">用户:{{book.user_id}}</p> 
	</div>
</div>
</template>

<script setup lang="ts">
import {defineProps} from 'vue'
const props = defineProps(['bookData'])
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