module.exports = {
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到MyBlog'
      return args
    })
  },
  productionSourceMap: false,
  devServer: {
    disableHostCheck: true  // 加上这个配置内网穿透
},
}
