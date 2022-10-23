<template>
    <BookList 
    :book-data="bookInfo"
    :isMyBookList="true"
    />
</template>
  
  <script setup lang="ts">
  import BookList from "../components/BookList.vue"
  import { useGetBookList } from '../utils/GetBookList'
  import { useMainStore } from '@/stores/UserState'

  const userState = useMainStore()
  let bookInfo = useGetBookList(userState.username)
  </script>
  
  <style>
  
  </style>