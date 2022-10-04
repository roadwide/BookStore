<template>
<div>
  <nav>
    <ul>
      <li><router-link to="/home">Home</router-link></li>
      <li><router-link to="/upload">Upload</router-link></li>
      <li><router-link to="/books">Books</router-link></li>
      <!-- 下面两个float:right 根据代码的顺序 从右往左排列，也即最右边是register -->
      <li style="float:right"><router-link to="/register" class="active">Register</router-link></li>
      <li style="float:right"><router-link to="/login">Login</router-link></li>

    </ul>
  </nav>
  <el-row align="middle" type="flex">

    <el-col :span="8" :offset="8">
      <router-view/>
    </el-col>

  </el-row>
</div>
</template>

<style scoped>
ul {
  list-style-type: none;    /* - 删除项目符号。导航条不需要列表项标记。*/
  margin: 0;   /* 设置 margin: 0; 和 padding: 0; 删除浏览器的默认设置。 */
  padding: 0;
  /* overflow:hidden 被添加到 ul 元素，以防止 li 元素超出列表。  没有这一句就没有背景色，不知道为什么*/
  overflow: hidden;
  border: 1px solid #e7e7e7;
  background-color: #f3f3f3;    /* 背景色的宽度跟屏幕一样，需要将背景色设置添加到ul元素中 */
}

li {
  /* -默认情况下，<li> 元素是块元素。在这里，我们删除每个列表项之前和之后的换行符，这样它们才能显示在一行。 */
  /* display: inline;    */
  float: left;   /* a标签 display: block;会导致导航栏变成纵向的，加上这一行恢复横向 */
  border-right: 1px solid #bbb;  /* 为所有列表项添加灰色右边框，最后一项（last-child）除外 */
}

li:last-child {
  border-right: none;
}

li a {
  display: block;   /* - 将链接显示为块元素可以使整个链接区域都可单击（不仅是文本），而且允许我们指定填充（如果需要，还可以指定高度，宽度，边距等） */
  color: #666;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  /* 背景色宽度跟导航栏的宽度一样   */
  /* background-color: #dddddd;    */
}

/* 当鼠标悬停时把链接颜色更改为 #111（黑色） */
li a:hover:not(.active) {
  background-color: #ddd;
}

/* class 向当前链接添加 "active" 类，这样用户就知道他/她在哪个页面上： */
/* li a.active表示 <li>标签里的<a>标签的class="active"才有效，不在li中的a标签无效 */
  li a.active {    
  color: white;
  background-color: #4CAF50;
}

</style>
