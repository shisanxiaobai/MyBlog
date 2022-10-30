<template>
  <div class="container" style="margin-top: 100px">
    <h2 class="text-center">
      联系方式
    </h2>
    <div class="grid-3_xs-1_sm-2_md-2" v-if="contacts.length">
      <div v-for="(contact, index) in contacts" :key="index" class="col">
        <a :href="contact.link">
          <vs-card type="3" class="center">
            <template #title>
              <h3><i :class="contact.icon" aria-hidden="true"></i>&nbsp;{{ contact.name }}</h3>
            </template>
            <template #text>
              <p>
                {{ contact.value }}
              </p>
            </template>
          </vs-card>
        </a>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Contacts',
  data() {
    return {
      contacts: [
        {
          name: 'QQ',
          link: 'tencent://message/?uin=1654622146&Site=&Menu=yes',
          value: '2221472921',
          icon: 'fa fa-qq'
        },
        {
          name: '微信',
          link: 'https://j.jcccj.cn/cFIZcj',
          value: 'Mgg8108',
          icon: 'fa fa-weixin'
        },
        {
          name: '邮箱',
          link: 'https://mail.163.com',
          value: 'lmy521777@163.com',
          icon: 'fa fa-envelope'
        }
      ]
    };
  },
  components: {
  }
};
</script>
