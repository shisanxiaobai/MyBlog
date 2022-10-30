<template>
  <div class="footer">
    <div class="container text-center">
      <h3>十三小白的博客</h3>
      <br />
      <p v-if="curYear === 2022">
        &copy; {{ curYear }} .&nbsp;
        前台页面参考于网上开源项目(CSDN、B站、github等).
      </p>
      <p v-else>
        &copy; 2022 - {{ curYear }} {{ config.name }}.&nbsp;
        Powered by github、CSDN，bolibili and so on.
      </p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Footer',
  data: function () {
    return {
      curYear: new Date().getFullYear()
    }
  },
  methods: {
    getYear: function () {
      return new Date().getFullYear()
    }
  }
}
</script>
