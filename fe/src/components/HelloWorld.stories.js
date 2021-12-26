import HelloWorld from '@/components/HelloWorld.vue';
export default {
  title: 'Example/HelloWorld',
  component: HelloWorld,
};

export const Template = (args, { argTypes }) => ({
  props: Object.keys(argTypes),
  components: { HelloWorld },
  template: `<HelloWorld v-bind="$props" title="HelloWorld"/>`,
});

export const Hello = Template.bind({});
Hello.args = {};
