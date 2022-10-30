<template>
  <div class="container text-center">
    <h1>
      关于我
    </h1>
    <hr />
    <h3>大家好我是十三小白，一个热爱it行业的普通人。</h3>
    <h3>20年本科毕业，目前正在学习go，打算转成golang开发。</h3>
    <h3>热爱篮球，爱好兴趣广泛，音乐、健身、街舞、钢琴、写作、历史，游戏等等，十八般武艺样样稀松。</h3>
    <h3>这个博客是为了学习go前后端分离开发而创作的入门项目，如果大家觉得对自己的学习，简历有帮助，欢迎大家学习研究。</h3>
    <h3>取自开源，出自开源，我会秉持这一理念去学习创作。愿大家未来可期，不忘初心，方得始终。</h3>
  </div>
</template>

<script>
export default {
  name: 'AboutMe',
  components: {
  },
  mounted: function () {
    this.changeTitle('About')
  }
}
</script>
