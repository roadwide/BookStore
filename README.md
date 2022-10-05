# BookStore

Test git branch  
https://segmentfault.com/a/1190000019714354

TODO:  
把所有登录状态判断写到App.vue中，其余组件都使用props传递的参数来判断是否登录  
把UserHome.vue 的 if (localStorage.getItem("userInfo") === null) 转移到App.vue中  

删掉UserHome这个组件，因为这个组件目前只用来显示用户名。  
将用户名显示在右上角