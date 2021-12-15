<template>
  <div>
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view />
  </div>
</template>
<style lang="scss">
#app {
  font-family: $body-font-family;
  text-align: center;
  max-width: 600px;
  width: 100%;
  height: 100vh;
  margin: 0 auto;
  background: {
    color: $bg-color;
  }
  font-size: $font-medium;
  color: $white-color;
}
</style>
