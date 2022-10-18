<template>
    <BookList 
    :book-data="bookInfo"
    :isMyBookList="false"
    />
</template>
  
  <script setup lang="ts">
  import BookList from "../components/BookList.vue"
  import { useGetBookList } from '../utils/GetBookList'
  let bookInfo = useGetBookList()

  </script>
  
  <style>
  
  </style>