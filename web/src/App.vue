<template>
  <nav>
    <router-link to="/home">Home</router-link> |
    <router-link to="/register">Register</router-link> |
    <router-link to="/login">Login</router-link> |
    <router-link to="/upload">Upload</router-link> |
    <router-link to="/books">Books</router-link>
  </nav>
  <el-row align="middle" type="flex">

    <el-col :span="8" :offset="8">
      <router-view/>
    </el-col>

  </el-row>
</template>

<style>

</style>
