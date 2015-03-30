## 1．创建型模式
前面讲过，社会化的分工越来越细，自然在软件设计方面也是如此，因此对象的创建和对象的使用分开也就成为了必然趋势。因为对象的创建会消耗掉系统的很多资源，所以单独对对象的创建进行研究，从而能够高效地创建对象就是创建型模式要探讨的问题。这里有6个具体的创建型模式可供研究，它们分别是：
- 工厂方法模式（Factory Method）；
- 抽象工厂模式（Abstract Factory）；
	- 提供一个创建一系列相关或相互依赖对象的接口，而无须指定他们具体的类。
	- 工厂方法把生产产品的方式封装起来了，但是一个工厂只能生产一类对象，当一个工厂需要生产多类产品的时候，就需要使用抽象工厂了。
- 创建者模式（Builder）；
	- 在软件系统设计中，有时候面临着一个“复杂系统”的创建工作，该对象通常由各个部分的子对象用一定的算法构成，或者说按一定的步骤组合而成；这些的算法和步骤是稳定的，而构成这个对象的子对象却经常由于需求改变而发生变化。
	- 生活中有许多这方面的例子，譬如安装一台电脑，他的组装过程基本上是不变的，都可以由主板，CPU，内存等按照某个稳定方式组合而成。然而，主板、CPU和内存等零件，本身都是可能多变的，这就适用于Builder模式。 	
- 原型模式（Prototype）；
- 单例模式（Singleton）。
	- 单例模式是一种常用的软件设计模式。在它的核心结构中只包含一个被称为单例类的特殊类。通过单例模式可以保证系统中一个类只有一个实例而且该实例易于外界访问，从而方便对实例个数的控制并节约系统资源。

## 2．结构型模式
在解决了对象的创建问题之后，对象的组成以及对象之间的依赖关系就成了开发人员关注的焦点，因为如何设计对象的结构、继承和依赖关系会影响到后续程序的维护性、代码的健壮性、耦合性等。对象结构的设计很容易体现出设计人员水平的高低，这里有7个具体的结构型模式可供研究，它们分别是：
- 外观模式（Facade）；
- 适配器模式（Adapter）；
- 代理模式（Proxy）；
- 装饰模式（Decorator）；
- 桥模式（Bridge）；
- 组合模式（Composite）；
- 享元模式（Flyweight）。
	- 如果个应用程序使用了大量的对象，而这些对象造成了**很大的存储开销**并且这些**对象都可以归纳成有限的种类**的时候就可以考虑是否可以使用享元模式。

## 3．行为型模式
在对象的结构和对象的创建问题都解决了之后，就剩下对象的行为问题了，如果对象的行为设计的好，那么对象的行为就会更清晰，它们之间的协作效率就会提高，这里有11个具体的行为型模式可供研究，它们分别是：
- 模板方法模式（Template Method）；
- 观察者模式（Observer）；
	-  在软件构建 过程中，我们需要为某些对象建立一种“通知依赖关系” --一个对象（目标对象）的状态发生改变，所有的依赖对象（观察者对象）都将得到通知 
- 状态模式（State）；
- 策略模式（Strategy）；
- 职责链模式（Chain of Responsibility）；
	- 使多个对象都有机会处理请求，从而避免请求的发送者和接受者之间的耦合关系。将这个对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理他为止。
- 命令模式（Command）；
- 访问者模式（Visitor）；
- 调停者模式（Mediator）；
- 备忘录模式（Memento）；
- 迭代器模式（Iterator）；
- 解释器模式（Interpreter）；


- - -
感谢 https://github.com/monochromegane/go_design_pattern

