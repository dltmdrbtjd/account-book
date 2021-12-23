import HelloWorld from '@/components/HelloWorld.vue';
export default {
  title: 'Example/HelloWorld',
  component: HelloWorld,
};

const Template = () => ({
  components: { HelloWorld },
  template: `<HelloWorld />`,
});

export const Hello = Template.bind({});
Hello.args = {};
