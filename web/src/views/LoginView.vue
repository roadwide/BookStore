<template>
  <!-- 
    这里如果 <UserLogin/>外面包一层div的话，父组件，也就是App.vue的router-view传的方法就过不来了
    但是props能够正确读取
    这里有点乱了，router-view到底是父组件，还是爷组件
    
    如果template中有多个根结点，不只是这个view里，view里的component也是
    会报错：Extraneous non-props attributes (class) were passed to component but could not be automatically i...
    原因是有些组件有多个根节点，在组件上添加属性，没有办法继承
    https://www.jianshu.com/p/4d19dd5ca4f5
    https://vuejs.org/guide/components/attrs.html

    所以，我的猜测是，对组件绑定的属性（方法或变量），这些组件必须是根节点才能继承。
    变量，在有多个根节点的情况下可以读取
    方法，在有多个根节点的情况下不能调用
    多个根节点时会报错
   -->
  <UserLogin/> 
</template>

<script>
import UserLogin from '../components/UserLogin.vue'
export default {
    components :{UserLogin}

}
</script>

<style>

</style>