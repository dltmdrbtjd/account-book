<template>
  <div>
    <button class="button">{{ button.title }}</button>
  </div>
</template>
<script>
import {ref, reactive} from 'vue';
export default {
  setup() {
    const button = reactive({title: 'Hello World!'});
    const button2 = ref({title: 'Hello World2'});

    return {button, button2};
  },
};
</script>
<style lang="scss" scoped>
.button {
  background-color: $light-gray-color;
  width: 200px;
  height: 200px;
  color: $white-color;
}
</style>
