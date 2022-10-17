<template>
    <BookList :book-data="bookInfo"/>
</template>
  
  <script setup lang="ts">
  import { defineProps}  from 'vue'
  import BookList from "../components/BookList.vue"
  import { useGetBookList } from '../utils/GetBookList'
  const props = defineProps(['userName'])
  let bookInfo = useGetBookList(props.userName)
  </script>
  
  <style>
  
  </style>