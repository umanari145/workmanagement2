<template>
  <div class="about">
    aaaaajojojojo
  </div>
</template>
<script>
export default {
  name: 'Reserve',
  data() {
    return {
      "work_room":"1",
      "button_message": "",
      "work_state":"1"
    };
  },
  methods:{

  },
  computed:{
    master_list:{
      get() {
        return this.$store.getters["master_list/get_master_list"];
      }
    }
  },
  mounted() {
  },
  created() {
    this.button_message = (this.work_state === "1") ? "保存します。" : "終了します";
  }
}
</script>
