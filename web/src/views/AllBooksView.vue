<template>
    <BookList :book-data="bookInfo"/>
</template>
  
  <script setup lang="ts">
  import BookList from "../components/BookList.vue"
  import { useGetBookList } from '../utils/GetBookList'
  let bookInfo = useGetBookList()
  </script>
  
  <style>
  
  </style>