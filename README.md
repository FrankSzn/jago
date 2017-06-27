# Jago

A simple Java interpreter written with Go language. This project is to learn JVM specification in depth and try to understand the
under-the-water mechanism when a Java program runs.

Basically, I only refer to "Java Virtual Machine Specification" and then compose how we should design one. 
I have to say many implementation details are naive and I make it as simplified as possible so as to
demonstrate the idea. For the educational purpose, it is enough.

# Roadmap

- Java class parser
- Interpreter
- Bridge JNI native method
- NO GC
- NO JIT
- NO multi-thread support

# Details

## trace classload and bytecode execution

```
GOROOT=/usr/local/Cellar/go/1.8.3/libexec
GOPATH=/Users/Chao/GoglandProjects/jago
/usr/local/Cellar/go/1.8.3/libexec/bin/go build -i -o /private/var/folders/84/7_y8dwv57nv49t5zn0dw9h5h0000gp/T/Build_main_go_and_rungo /Users/Chao/GoglandProjects/jago/main.go
/private/var/folders/84/7_y8dwv57nv49t5zn0dw9h5h0000gp/T/Build_main_go_and_rungo
	FDFD

  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ Main 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ java/lang/Object 
      ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
      ↳ java/lang/Class 
        ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
        ↳ java/io/Serializable 
        ↱ java/io/Serializable 
        ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
        ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
        ↳ java/lang/reflect/GenericDeclaration 
          ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
          ↳ java/lang/reflect/AnnotatedElement 
          ↱ java/lang/reflect/AnnotatedElement 
          ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
        ↱ java/lang/reflect/GenericDeclaration 
        ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
        ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
        ↳ java/lang/reflect/Type 
        ↱ java/lang/reflect/Type 
        ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
      ↱ java/lang/Class 
      ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↱ java/lang/Object 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↱ Main 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
🍷 Main.<clinit>()V 
   0000 ➢ iconst_2          
   0001 ➢ putstatic          #22
   0004 ➢ return            
🍷 java/lang/Object.<clinit>()V 
   0000 ➢ invokestatic       #16
🍺invoke native method java/lang/Object.registerNatives()V 😎

   0003 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0000 ➢ iconst_2          
   0001 ➢ istore_1          
   0002 ➢ iconst_2          
   0003 ➢ istore_2          
   0004 ➢ iload_1           
   0005 ➢ iload_2           
   0006 ➢ iadd              
   0007 ➢ istore_3          
   0008 ➢ iload_3           
   0009 ➢ invokestatic       #2
🍷 Main.increase(I)I 
   0000 ➢ iload_0           
   0001 ➢ iconst_1          
   0002 ➢ iadd              
   0003 ➢ ireturn           
🍷 Main.main([Ljava/lang/String;)V 
   0012 ➢ istore             #4
   0014 ➢ new                #3
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ B 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ A 
    ↱ A 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↱ B 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0017 ➢ dup               
   0018 ➢ invokespecial      #4
🍷 B.<init>()V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 A.<init>()V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 A.<init>()V 
   0004 ➢ return            
🍷 B.<init>()V 
   0004 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0021 ➢ astore             #5
   0023 ➢ iconst_1          
   0024 ➢ putstatic          #5
   0027 ➢ aload              #5
   0029 ➢ iconst_2          
   0030 ➢ putfield           #6
   0033 ➢ aload              #5
   0035 ➢ iconst_3          
   0036 ➢ putfield           #7
   0039 ➢ aload              #5
   0041 ➢ invokevirtual      #8
🍷 B.foo()V 
   0000 ➢ aload_0           
   0001 ➢ getfield           #2
   0004 ➢ istore_1          
   0005 ➢ getstatic          #3
   0008 ➢ istore_2          
   0009 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0044 ➢ iload_1           
   0045 ➢ i2d               
   0046 ➢ ldc2_w             #9
   0049 ➢ ddiv              
   0050 ➢ dstore             #6
   0052 ➢ new                #11
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ GVM 
  ↱ GVM 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0055 ➢ dup               
   0056 ➢ invokespecial      #12
🍷 GVM.<init>()V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 GVM.<init>()V 
   0004 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0059 ➢ astore             #8
   0061 ➢ aload              #8
   0063 ➢ ldc                #13
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/String 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ java/lang/Comparable 
    ↱ java/lang/Comparable 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ java/lang/CharSequence 
    ↱ java/lang/CharSequence 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↱ java/lang/String 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ [C 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ java/lang/Cloneable 
    ↱ java/lang/Cloneable 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↱ [C 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0065 ➢ invokestatic       #14
🍷 GVM.<clinit>()V 
   0000 ➢ bipush            
   0002 ➢ putstatic          #4
   0005 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0065 ➢ invokestatic       #14
🍺invoke native method GVM.toUpper(Ljava/lang/String;)Ljava/lang/String; 😎

   0068 ➢ invokevirtual      #15
🍷 GVM.println(Ljava/lang/Object;)V 
   0000 ➢ aload_0           
   0001 ➢ aload_1           
   0002 ➢ invokestatic       #2
🍷 java/lang/String.<clinit>()V 
   0000 ➢ iconst_0          
   0001 ➢ anewarray          #143
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/io/ObjectStreamField 
  ↱ java/io/ObjectStreamField 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ [Ljava/io/ObjectStreamField; 
  ↱ [Ljava/io/ObjectStreamField; 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0004 ➢ putstatic          #144
   0007 ➢ new                #145
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/String$CaseInsensitiveComparator 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ java/util/Comparator 
    ↱ java/util/Comparator 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↱ java/lang/String$CaseInsensitiveComparator 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0010 ➢ dup               
   0011 ➢ aconst_null       
   0012 ➢ invokespecial      #146
🍷 java/lang/String$CaseInsensitiveComparator.<init>(Ljava/lang/String$1;)V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 java/lang/String$CaseInsensitiveComparator.<init>()V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #2
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 java/lang/String$CaseInsensitiveComparator.<init>()V 
   0004 ➢ return            
🍷 java/lang/String$CaseInsensitiveComparator.<init>(Ljava/lang/String$1;)V 
   0004 ➢ return            
🍷 java/lang/String.<clinit>()V 
   0015 ➢ putstatic          #55
   0018 ➢ return            
🍷 GVM.println(Ljava/lang/Object;)V 
   0002 ➢ invokestatic       #2
🍷 java/lang/String.valueOf(Ljava/lang/Object;)Ljava/lang/String; 
   0000 ➢ aload_0           
   0001 ➢ ifnonnull         
   0009 ➢ aload_0           
   0010 ➢ invokevirtual      #135
🍷 java/lang/String.toString()Ljava/lang/String; 
   0000 ➢ aload_0           
   0001 ➢ areturn           
🍷 java/lang/String.valueOf(Ljava/lang/Object;)Ljava/lang/String; 
   0013 ➢ areturn           
🍷 GVM.println(Ljava/lang/Object;)V 
   0005 ➢ invokespecial      #3
🍺invoke native method GVM.println(Ljava/lang/String;)V 😎

   0008 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0071 ➢ bipush            
   0073 ➢ anewarray          #16
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ C 
  ↱ C 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ [LC; 
  ↱ [LC; 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0076 ➢ astore             #9
   0078 ➢ ldc                #17
   0080 ➢ istore             #10
   0082 ➢ iload              #10
   0084 ➢ i2c               
   0085 ➢ istore             #11
   0087 ➢ aload              #8
   0089 ➢ iload              #11
   0091 ➢ invokestatic       #18
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/Character 
  ↱ java/lang/Character 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

🍷 java/lang/Character.<clinit>()V 
   0000 ➢ ldc                #1
   0002 ➢ invokevirtual      #102
🍷 java/lang/Class.desiredAssertionStatus()Z 
   0000 ➢ aload_0           
   0001 ➢ invokevirtual      #257
🍷 java/lang/Class.getClassLoader()Ljava/lang/ClassLoader; 
   0000 ➢ aload_0           
   0001 ➢ invokevirtual      #67
🍷 java/lang/Class.getClassLoader0()Ljava/lang/ClassLoader; 
   0000 ➢ aload_0           
   0001 ➢ getfield           #10
   0004 ➢ areturn           
🍷 java/lang/Class.getClassLoader()Ljava/lang/ClassLoader; 
   0004 ➢ astore_1          
   0005 ➢ aload_1           
   0006 ➢ ifnonnull         
   0009 ➢ aconst_null       
   0010 ➢ areturn           
🍷 java/lang/Class.desiredAssertionStatus()Z 
   0004 ➢ astore_1          
   0005 ➢ aload_1           
   0006 ➢ ifnonnull         
   0009 ➢ aload_0           
   0010 ➢ invokestatic       #258
🍷 java/lang/Class.<clinit>()V 
   0000 ➢ invokestatic       #330
🍺invoke native method java/lang/Class.registerNatives()V 😎

   0003 ➢ iconst_1          
   0004 ➢ putstatic          #2
   0007 ➢ iconst_0          
   0008 ➢ anewarray          #331
   0011 ➢ putstatic          #332
   0014 ➢ iconst_0          
   0015 ➢ putstatic          #1
   0018 ➢ return            
🍷 java/lang/Class.desiredAssertionStatus()Z 
   0010 ➢ invokestatic       #258
🍺invoke native method java/lang/Class.desiredAssertionStatus0(Ljava/lang/Class;)Z 😎

   0013 ➢ ireturn           
🍷 java/lang/Character.<clinit>()V 
   0005 ➢ ifne              
   0012 ➢ iconst_0          
   0013 ➢ putstatic          #83
   0016 ➢ ldc                #103
   0018 ➢ invokestatic       #104
🍺invoke native method java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class; 😎

   0021 ➢ putstatic          #105
   0024 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0091 ➢ invokestatic       #18
🍷 java/lang/Character.valueOf(C)Ljava/lang/Character; 
   0000 ➢ iload_0           
   0001 ➢ bipush            
   0003 ➢ if_icmpgt         
   0012 ➢ new                #1
   0015 ➢ dup               
   0016 ➢ iload_0           
   0017 ➢ invokespecial      #6
🍷 java/lang/Character.<init>(C)V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #3
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 java/lang/Character.<init>(C)V 
   0004 ➢ aload_0           
   0005 ➢ iload_1           
   0006 ➢ putfield           #4
   0009 ➢ return            
🍷 java/lang/Character.valueOf(C)Ljava/lang/Character; 
   0020 ➢ areturn           
🍷 Main.main([Ljava/lang/String;)V 
   0094 ➢ invokevirtual      #15
🍷 GVM.println(Ljava/lang/Object;)V 
   0000 ➢ aload_0           
   0001 ➢ aload_1           
   0002 ➢ invokestatic       #2
🍷 java/lang/String.valueOf(Ljava/lang/Object;)Ljava/lang/String; 
   0000 ➢ aload_0           
   0001 ➢ ifnonnull         
   0009 ➢ aload_0           
   0010 ➢ invokevirtual      #135
🍷 java/lang/Character.toString()Ljava/lang/String; 
   0000 ➢ iconst_1          
   0001 ➢ newarray          
   0003 ➢ dup               
   0004 ➢ iconst_0          
   0005 ➢ aload_0           
   0006 ➢ getfield           #4
   0009 ➢ castore           
   0010 ➢ astore_1          
   0011 ➢ aload_1           
   0012 ➢ invokestatic       #9
🍷 java/lang/String.valueOf([C)Ljava/lang/String; 
   0000 ➢ new                #43
   0003 ➢ dup               
   0004 ➢ aload_0           
   0005 ➢ invokespecial      #136
🍷 java/lang/String.<init>([C)V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 java/lang/String.<init>([C)V 
   0004 ➢ aload_0           
   0005 ➢ aload_1           
   0006 ➢ aload_1           
   0007 ➢ arraylength       
   0008 ➢ invokestatic       #5
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/util/Arrays 
  ↱ java/util/Arrays 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

🍷 java/util/Arrays.<clinit>()V 
   0000 ➢ ldc                #22
   0002 ➢ invokevirtual      #197
🍷 java/lang/Class.desiredAssertionStatus()Z 
   0000 ➢ aload_0           
   0001 ➢ invokevirtual      #257
🍷 java/lang/Class.getClassLoader()Ljava/lang/ClassLoader; 
   0000 ➢ aload_0           
   0001 ➢ invokevirtual      #67
🍷 java/lang/Class.getClassLoader0()Ljava/lang/ClassLoader; 
   0000 ➢ aload_0           
   0001 ➢ getfield           #10
   0004 ➢ areturn           
🍷 java/lang/Class.getClassLoader()Ljava/lang/ClassLoader; 
   0004 ➢ astore_1          
   0005 ➢ aload_1           
   0006 ➢ ifnonnull         
   0009 ➢ aconst_null       
   0010 ➢ areturn           
🍷 java/lang/Class.desiredAssertionStatus()Z 
   0004 ➢ astore_1          
   0005 ➢ aload_1           
   0006 ➢ ifnonnull         
   0009 ➢ aload_0           
   0010 ➢ invokestatic       #258
🍺invoke native method java/lang/Class.desiredAssertionStatus0(Ljava/lang/Class;)Z 😎

   0013 ➢ ireturn           
🍷 java/util/Arrays.<clinit>()V 
   0005 ➢ ifne              
   0012 ➢ iconst_0          
   0013 ➢ putstatic          #124
   0016 ➢ return            
🍷 java/lang/String.<init>([C)V 
   0008 ➢ invokestatic       #5
🍷 java/util/Arrays.copyOf([CI)[C 
   0000 ➢ iload_1           
   0001 ➢ newarray          
   0003 ➢ astore_2          
   0004 ➢ aload_0           
   0005 ➢ iconst_0          
   0006 ➢ aload_2           
   0007 ➢ iconst_0          
   0008 ➢ aload_0           
   0009 ➢ arraylength       
   0010 ➢ iload_1           
   0011 ➢ invokestatic       #100
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/Math 
  ↱ java/lang/Math 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

🍷 java/lang/Math.<clinit>()V 
   0000 ➢ ldc                #10
   0002 ➢ invokevirtual      #108
🍷 java/lang/Class.desiredAssertionStatus()Z 
   0000 ➢ aload_0           
   0001 ➢ invokevirtual      #257
🍷 java/lang/Class.getClassLoader()Ljava/lang/ClassLoader; 
   0000 ➢ aload_0           
   0001 ➢ invokevirtual      #67
🍷 java/lang/Class.getClassLoader0()Ljava/lang/ClassLoader; 
   0000 ➢ aload_0           
   0001 ➢ getfield           #10
   0004 ➢ areturn           
🍷 java/lang/Class.getClassLoader()Ljava/lang/ClassLoader; 
   0004 ➢ astore_1          
   0005 ➢ aload_1           
   0006 ➢ ifnonnull         
   0009 ➢ aconst_null       
   0010 ➢ areturn           
🍷 java/lang/Class.desiredAssertionStatus()Z 
   0004 ➢ astore_1          
   0005 ➢ aload_1           
   0006 ➢ ifnonnull         
   0009 ➢ aload_0           
   0010 ➢ invokestatic       #258
🍺invoke native method java/lang/Class.desiredAssertionStatus0(Ljava/lang/Class;)Z 😎

   0013 ➢ ireturn           
🍷 java/lang/Math.<clinit>()V 
   0005 ➢ ifne              
   0012 ➢ iconst_0          
   0013 ➢ putstatic          #67
   0016 ➢ ldc                #109
   0018 ➢ invokestatic       #24
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/Float 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ java/lang/Number 
    ↱ java/lang/Number 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↱ java/lang/Float 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

🍷 java/lang/Float.<clinit>()V 
   0000 ➢ ldc                #34
   0002 ➢ invokestatic       #35
🍺invoke native method java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class; 😎

   0005 ➢ putstatic          #36
   0008 ➢ return            
🍷 java/lang/Math.<clinit>()V 
   0018 ➢ invokestatic       #24
🍺invoke native method java/lang/Float.floatToRawIntBits(F)I 😎

   0021 ➢ i2l               
   0022 ➢ putstatic          #60
   0025 ➢ ldc2_w             #110
   0028 ➢ invokestatic       #29
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/Double 
  ↱ java/lang/Double 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

🍷 java/lang/Double.<clinit>()V 
   0000 ➢ ldc                #61
   0002 ➢ invokestatic       #62
🍺invoke native method java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class; 😎

   0005 ➢ putstatic          #63
   0008 ➢ return            
🍷 java/lang/Math.<clinit>()V 
   0028 ➢ invokestatic       #29
🍺invoke native method java/lang/Double.doubleToRawLongBits(D)J 😎

   0031 ➢ putstatic          #61
   0034 ➢ sipush            
   0037 ➢ invokestatic       #70
🍷 java/lang/Math.powerOfTwoD(I)D 
   0000 ➢ getstatic          #67
   0003 ➢ ifne              
   0006 ➢ iload_0           
   0007 ➢ sipush            
   0010 ➢ if_icmplt         
   0013 ➢ iload_0           
   0014 ➢ sipush            
   0017 ➢ if_icmple         
   0028 ➢ iload_0           
   0029 ➢ i2l               
   0030 ➢ ldc2_w             #87
   0033 ➢ ladd              
   0034 ➢ bipush            
   0036 ➢ lshl              
   0037 ➢ ldc2_w             #31
   0040 ➢ land              
   0041 ➢ invokestatic       #71
🍺invoke native method java/lang/Double.longBitsToDouble(J)D 😎

   0044 ➢ dreturn           
🍷 java/lang/Math.<clinit>()V 
   0040 ➢ putstatic          #107
   0043 ➢ sipush            
   0046 ➢ invokestatic       #70
🍷 java/lang/Math.powerOfTwoD(I)D 
   0000 ➢ getstatic          #67
   0003 ➢ ifne              
   0006 ➢ iload_0           
   0007 ➢ sipush            
   0010 ➢ if_icmplt         
   0013 ➢ iload_0           
   0014 ➢ sipush            
   0017 ➢ if_icmple         
   0028 ➢ iload_0           
   0029 ➢ i2l               
   0030 ➢ ldc2_w             #87
   0033 ➢ ladd              
   0034 ➢ bipush            
   0036 ➢ lshl              
   0037 ➢ ldc2_w             #31
   0040 ➢ land              
   0041 ➢ invokestatic       #71
🍺invoke native method java/lang/Double.longBitsToDouble(J)D 😎

   0044 ➢ dreturn           
🍷 java/lang/Math.<clinit>()V 
   0049 ➢ putstatic          #105
   0052 ➢ return            
🍷 java/util/Arrays.copyOf([CI)[C 
   0011 ➢ invokestatic       #100
🍷 java/lang/Math.min(II)I 
   0000 ➢ iload_0           
   0001 ➢ iload_1           
   0002 ➢ if_icmpgt         
   0005 ➢ iload_0           
   0006 ➢ goto              
   0010 ➢ ireturn           
🍷 java/util/Arrays.copyOf([CI)[C 
   0014 ➢ invokestatic       #65
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/System 
豊
  ↱ java/lang/System 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

🍷 java/lang/System.<clinit>()V 
   0000 ➢ invokestatic       #102
🍺invoke native method java/lang/System.registerNatives()V 😎

   0003 ➢ aconst_null       
   0004 ➢ putstatic          #103
   0007 ➢ aconst_null       
   0008 ➢ putstatic          #104
   0011 ➢ aconst_null       
   0012 ➢ putstatic          #105
   0015 ➢ aconst_null       
   0016 ➢ putstatic          #27
   0019 ➢ aconst_null       
   0020 ➢ putstatic          #6
   0023 ➢ return            
🍷 java/util/Arrays.copyOf([CI)[C 
   0014 ➢ invokestatic       #65
🍺invoke native method java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V 😎

   0017 ➢ aload_2           
   0018 ➢ areturn           
🍷 java/lang/String.<init>([C)V 
   0011 ➢ putfield           #3
   0014 ➢ return            
🍷 java/lang/String.valueOf([C)Ljava/lang/String; 
   0008 ➢ areturn           
🍷 java/lang/Character.toString()Ljava/lang/String; 
   0015 ➢ areturn           
🍷 java/lang/String.valueOf(Ljava/lang/Object;)Ljava/lang/String; 
   0013 ➢ areturn           
🍷 GVM.println(Ljava/lang/Object;)V 
   0005 ➢ invokespecial      #3
🍺invoke native method GVM.println(Ljava/lang/String;)V 😎

   0008 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0097 ➢ aload              #8
   0099 ➢ new                #16
   0102 ➢ dup               
   0103 ➢ invokespecial      #19
🍷 C.<init>()V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 C.<init>()V 
   0004 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0106 ➢ invokevirtual      #20
🍷 C.ss()Ljava/lang/String; 
   0000 ➢ iconst_1          
   0001 ➢ newarray          
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ [I 
  ↱ [I 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0003 ➢ dup               
   0004 ➢ iconst_0          
   0005 ➢ iconst_1          
   0006 ➢ iastore           
   0007 ➢ astore_1          
   0008 ➢ aload_1           
   0009 ➢ invokevirtual      #6
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0000 ➢ new                #1
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/StringBuilder 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↳ java/lang/AbstractStringBuilder 
      ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
      ↳ java/lang/Appendable 
      ↱ java/lang/Appendable 
      ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ↱ java/lang/AbstractStringBuilder 
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↱ java/lang/StringBuilder 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

   0003 ➢ dup               
   0004 ➢ invokespecial      #2
🍷 java/lang/StringBuilder.<init>()V 
   0000 ➢ aload_0           
   0001 ➢ bipush            
   0003 ➢ invokespecial      #1
🍷 java/lang/AbstractStringBuilder.<init>(I)V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #2
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 java/lang/AbstractStringBuilder.<init>(I)V 
   0004 ➢ aload_0           
   0005 ➢ iload_1           
   0006 ➢ newarray          
   0008 ➢ putfield           #3
   0011 ➢ return            
🍷 java/lang/StringBuilder.<init>()V 
   0006 ➢ return            
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0007 ➢ aload_0           
   0008 ➢ invokevirtual      #3
🍺invoke native method java/lang/Object.getClass()Ljava/lang/Class; 😎

   0011 ➢ invokevirtual      #4
🍷 java/lang/Class.getName()Ljava/lang/String; 
   0000 ➢ aload_0           
   0001 ➢ getfield           #65
   0004 ➢ astore_1          
   0005 ➢ aload_1           
   0006 ➢ ifnonnull         
   0009 ➢ aload_0           
   0010 ➢ aload_0           
   0011 ➢ invokespecial      #66
🍺invoke native method java/lang/Class.getName0()Ljava/lang/String; 😎

   0014 ➢ dup               
   0015 ➢ astore_1          
   0016 ➢ putfield           #65
   0019 ➢ aload_1           
   0020 ➢ areturn           
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0014 ➢ invokevirtual      #5
🍷 java/lang/StringBuilder.append(Ljava/lang/String;)Ljava/lang/StringBuilder; 
   0000 ➢ aload_0           
   0001 ➢ aload_1           
   0002 ➢ invokespecial      #8
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0000 ➢ aload_1           
   0001 ➢ ifnonnull         
   0009 ➢ aload_1           
   0010 ➢ invokevirtual      #29
🍷 java/lang/String.length()I 
   0000 ➢ aload_0           
   0001 ➢ getfield           #3
   0004 ➢ arraylength       
   0005 ➢ ireturn           
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0013 ➢ istore_2          
   0014 ➢ aload_0           
   0015 ➢ aload_0           
   0016 ➢ getfield           #4
   0019 ➢ iload_2           
   0020 ➢ iadd              
   0021 ➢ invokespecial      #5
🍷 java/lang/AbstractStringBuilder.ensureCapacityInternal(I)V 
   0000 ➢ iload_1           
   0001 ➢ aload_0           
   0002 ➢ getfield           #3
   0005 ➢ arraylength       
   0006 ➢ isub              
   0007 ➢ ifle              
   0026 ➢ return            
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0024 ➢ aload_1           
   0025 ➢ iconst_0          
   0026 ➢ iload_2           
   0027 ➢ aload_0           
   0028 ➢ getfield           #3
   0031 ➢ aload_0           
   0032 ➢ getfield           #4
   0035 ➢ invokevirtual      #30
🍷 java/lang/String.getChars(II[CI)V 
   0000 ➢ iload_1           
   0001 ➢ ifge              
   0013 ➢ iload_2           
   0014 ➢ aload_0           
   0015 ➢ getfield           #3
   0018 ➢ arraylength       
   0019 ➢ if_icmple         
   0031 ➢ iload_1           
   0032 ➢ iload_2           
   0033 ➢ if_icmple         
   0047 ➢ aload_0           
   0048 ➢ getfield           #3
   0051 ➢ iload_1           
   0052 ➢ aload_3           
   0053 ➢ iload              #4
   0055 ➢ iload_2           
   0056 ➢ iload_1           
   0057 ➢ isub              
   0058 ➢ invokestatic       #37
🍺invoke native method java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V 😎

   0061 ➢ return            
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0038 ➢ aload_0           
   0039 ➢ dup               
   0040 ➢ getfield           #4
   0043 ➢ iload_2           
   0044 ➢ iadd              
   0045 ➢ putfield           #4
   0048 ➢ aload_0           
   0049 ➢ areturn           
🍷 java/lang/StringBuilder.append(Ljava/lang/String;)Ljava/lang/StringBuilder; 
   0005 ➢ pop               
   0006 ➢ aload_0           
   0007 ➢ areturn           
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0017 ➢ ldc                #6
   0019 ➢ invokevirtual      #5
🍷 java/lang/StringBuilder.append(Ljava/lang/String;)Ljava/lang/StringBuilder; 
   0000 ➢ aload_0           
   0001 ➢ aload_1           
   0002 ➢ invokespecial      #8
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0000 ➢ aload_1           
   0001 ➢ ifnonnull         
   0009 ➢ aload_1           
   0010 ➢ invokevirtual      #29
🍷 java/lang/String.length()I 
   0000 ➢ aload_0           
   0001 ➢ getfield           #3
   0004 ➢ arraylength       
   0005 ➢ ireturn           
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0013 ➢ istore_2          
   0014 ➢ aload_0           
   0015 ➢ aload_0           
   0016 ➢ getfield           #4
   0019 ➢ iload_2           
   0020 ➢ iadd              
   0021 ➢ invokespecial      #5
🍷 java/lang/AbstractStringBuilder.ensureCapacityInternal(I)V 
   0000 ➢ iload_1           
   0001 ➢ aload_0           
   0002 ➢ getfield           #3
   0005 ➢ arraylength       
   0006 ➢ isub              
   0007 ➢ ifle              
   0026 ➢ return            
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0024 ➢ aload_1           
   0025 ➢ iconst_0          
   0026 ➢ iload_2           
   0027 ➢ aload_0           
   0028 ➢ getfield           #3
   0031 ➢ aload_0           
   0032 ➢ getfield           #4
   0035 ➢ invokevirtual      #30
🍷 java/lang/String.getChars(II[CI)V 
   0000 ➢ iload_1           
   0001 ➢ ifge              
   0013 ➢ iload_2           
   0014 ➢ aload_0           
   0015 ➢ getfield           #3
   0018 ➢ arraylength       
   0019 ➢ if_icmple         
   0031 ➢ iload_1           
   0032 ➢ iload_2           
   0033 ➢ if_icmple         
   0047 ➢ aload_0           
   0048 ➢ getfield           #3
   0051 ➢ iload_1           
   0052 ➢ aload_3           
   0053 ➢ iload              #4
   0055 ➢ iload_2           
   0056 ➢ iload_1           
   0057 ➢ isub              
   0058 ➢ invokestatic       #37
🍺invoke native method java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V 😎

   0061 ➢ return            
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0038 ➢ aload_0           
   0039 ➢ dup               
   0040 ➢ getfield           #4
   0043 ➢ iload_2           
   0044 ➢ iadd              
   0045 ➢ putfield           #4
   0048 ➢ aload_0           
   0049 ➢ areturn           
🍷 java/lang/StringBuilder.append(Ljava/lang/String;)Ljava/lang/StringBuilder; 
   0005 ➢ pop               
   0006 ➢ aload_0           
   0007 ➢ areturn           
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0022 ➢ aload_0           
   0023 ➢ invokevirtual      #7
🍺invoke native method java/lang/Object.hashCode()I 😎

   0026 ➢ invokestatic       #8
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ↳ java/lang/Integer 
  ↱ java/lang/Integer 
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧

🍷 java/lang/Integer.<clinit>()V 
   0000 ➢ ldc                #85
   0002 ➢ invokestatic       #86
🍺invoke native method java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class; 😎

   0005 ➢ putstatic          #87
   0008 ➢ bipush            
   0010 ➢ newarray          
   0012 ➢ dup               
   0013 ➢ iconst_0          
   0014 ➢ bipush            
   0016 ➢ castore           
   0017 ➢ dup               
   0018 ➢ iconst_1          
   0019 ➢ bipush            
   0021 ➢ castore           
   0022 ➢ dup               
   0023 ➢ iconst_2          
   0024 ➢ bipush            
   0026 ➢ castore           
   0027 ➢ dup               
   0028 ➢ iconst_3          
   0029 ➢ bipush            
   0031 ➢ castore           
   0032 ➢ dup               
   0033 ➢ iconst_4          
   0034 ➢ bipush            
   0036 ➢ castore           
   0037 ➢ dup               
   0038 ➢ iconst_5          
   0039 ➢ bipush            
   0041 ➢ castore           
   0042 ➢ dup               
   0043 ➢ bipush            
   0045 ➢ bipush            
   0047 ➢ castore           
   0048 ➢ dup               
   0049 ➢ bipush            
   0051 ➢ bipush            
   0053 ➢ castore           
   0054 ➢ dup               
   0055 ➢ bipush            
   0057 ➢ bipush            
   0059 ➢ castore           
   0060 ➢ dup               
   0061 ➢ bipush            
   0063 ➢ bipush            
   0065 ➢ castore           
   0066 ➢ dup               
   0067 ➢ bipush            
   0069 ➢ bipush            
   0071 ➢ castore           
   0072 ➢ dup               
   0073 ➢ bipush            
   0075 ➢ bipush            
   0077 ➢ castore           
   0078 ➢ dup               
   0079 ➢ bipush            
   0081 ➢ bipush            
   0083 ➢ castore           
   0084 ➢ dup               
   0085 ➢ bipush            
   0087 ➢ bipush            
   0089 ➢ castore           
   0090 ➢ dup               
   0091 ➢ bipush            
   0093 ➢ bipush            
   0095 ➢ castore           
   0096 ➢ dup               
   0097 ➢ bipush            
   0099 ➢ bipush            
   0101 ➢ castore           
   0102 ➢ dup               
   0103 ➢ bipush            
   0105 ➢ bipush            
   0107 ➢ castore           
   0108 ➢ dup               
   0109 ➢ bipush            
   0111 ➢ bipush            
   0113 ➢ castore           
   0114 ➢ dup               
   0115 ➢ bipush            
   0117 ➢ bipush            
   0119 ➢ castore           
   0120 ➢ dup               
   0121 ➢ bipush            
   0123 ➢ bipush            
   0125 ➢ castore           
   0126 ➢ dup               
   0127 ➢ bipush            
   0129 ➢ bipush            
   0131 ➢ castore           
   0132 ➢ dup               
   0133 ➢ bipush            
   0135 ➢ bipush            
   0137 ➢ castore           
   0138 ➢ dup               
   0139 ➢ bipush            
   0141 ➢ bipush            
   0143 ➢ castore           
   0144 ➢ dup               
   0145 ➢ bipush            
   0147 ➢ bipush            
   0149 ➢ castore           
   0150 ➢ dup               
   0151 ➢ bipush            
   0153 ➢ bipush            
   0155 ➢ castore           
   0156 ➢ dup               
   0157 ➢ bipush            
   0159 ➢ bipush            
   0161 ➢ castore           
   0162 ➢ dup               
   0163 ➢ bipush            
   0165 ➢ bipush            
   0167 ➢ castore           
   0168 ➢ dup               
   0169 ➢ bipush            
   0171 ➢ bipush            
   0173 ➢ castore           
   0174 ➢ dup               
   0175 ➢ bipush            
   0177 ➢ bipush            
   0179 ➢ castore           
   0180 ➢ dup               
   0181 ➢ bipush            
   0183 ➢ bipush            
   0185 ➢ castore           
   0186 ➢ dup               
   0187 ➢ bipush            
   0189 ➢ bipush            
   0191 ➢ castore           
   0192 ➢ dup               
   0193 ➢ bipush            
   0195 ➢ bipush            
   0197 ➢ castore           
   0198 ➢ dup               
   0199 ➢ bipush            
   0201 ➢ bipush            
   0203 ➢ castore           
   0204 ➢ dup               
   0205 ➢ bipush            
   0207 ➢ bipush            
   0209 ➢ castore           
   0210 ➢ dup               
   0211 ➢ bipush            
   0213 ➢ bipush            
   0215 ➢ castore           
   0216 ➢ dup               
   0217 ➢ bipush            
   0219 ➢ bipush            
   0221 ➢ castore           
   0222 ➢ putstatic          #5
   0225 ➢ bipush            
   0227 ➢ newarray          
   0229 ➢ dup               
   0230 ➢ iconst_0          
   0231 ➢ bipush            
   0233 ➢ castore           
   0234 ➢ dup               
   0235 ➢ iconst_1          
   0236 ➢ bipush            
   0238 ➢ castore           
   0239 ➢ dup               
   0240 ➢ iconst_2          
   0241 ➢ bipush            
   0243 ➢ castore           
   0244 ➢ dup               
   0245 ➢ iconst_3          
   0246 ➢ bipush            
   0248 ➢ castore           
   0249 ➢ dup               
   0250 ➢ iconst_4          
   0251 ➢ bipush            
   0253 ➢ castore           
   0254 ➢ dup               
   0255 ➢ iconst_5          
   0256 ➢ bipush            
   0258 ➢ castore           
   0259 ➢ dup               
   0260 ➢ bipush            
   0262 ➢ bipush            
   0264 ➢ castore           
   0265 ➢ dup               
   0266 ➢ bipush            
   0268 ➢ bipush            
   0270 ➢ castore           
   0271 ➢ dup               
   0272 ➢ bipush            
   0274 ➢ bipush            
   0276 ➢ castore           
   0277 ➢ dup               
   0278 ➢ bipush            
   0280 ➢ bipush            
   0282 ➢ castore           
   0283 ➢ dup               
   0284 ➢ bipush            
   0286 ➢ bipush            
   0288 ➢ castore           
   0289 ➢ dup               
   0290 ➢ bipush            
   0292 ➢ bipush            
   0294 ➢ castore           
   0295 ➢ dup               
   0296 ➢ bipush            
   0298 ➢ bipush            
   0300 ➢ castore           
   0301 ➢ dup               
   0302 ➢ bipush            
   0304 ➢ bipush            
   0306 ➢ castore           
   0307 ➢ dup               
   0308 ➢ bipush            
   0310 ➢ bipush            
   0312 ➢ castore           
   0313 ➢ dup               
   0314 ➢ bipush            
   0316 ➢ bipush            
   0318 ➢ castore           
   0319 ➢ dup               
   0320 ➢ bipush            
   0322 ➢ bipush            
   0324 ➢ castore           
   0325 ➢ dup               
   0326 ➢ bipush            
   0328 ➢ bipush            
   0330 ➢ castore           
   0331 ➢ dup               
   0332 ➢ bipush            
   0334 ➢ bipush            
   0336 ➢ castore           
   0337 ➢ dup               
   0338 ➢ bipush            
   0340 ➢ bipush            
   0342 ➢ castore           
   0343 ➢ dup               
   0344 ➢ bipush            
   0346 ➢ bipush            
   0348 ➢ castore           
   0349 ➢ dup               
   0350 ➢ bipush            
   0352 ➢ bipush            
   0354 ➢ castore           
   0355 ➢ dup               
   0356 ➢ bipush            
   0358 ➢ bipush            
   0360 ➢ castore           
   0361 ➢ dup               
   0362 ➢ bipush            
   0364 ➢ bipush            
   0366 ➢ castore           
   0367 ➢ dup               
   0368 ➢ bipush            
   0370 ➢ bipush            
   0372 ➢ castore           
   0373 ➢ dup               
   0374 ➢ bipush            
   0376 ➢ bipush            
   0378 ➢ castore           
   0379 ➢ dup               
   0380 ➢ bipush            
   0382 ➢ bipush            
   0384 ➢ castore           
   0385 ➢ dup               
   0386 ➢ bipush            
   0388 ➢ bipush            
   0390 ➢ castore           
   0391 ➢ dup               
   0392 ➢ bipush            
   0394 ➢ bipush            
   0396 ➢ castore           
   0397 ➢ dup               
   0398 ➢ bipush            
   0400 ➢ bipush            
   0402 ➢ castore           
   0403 ➢ dup               
   0404 ➢ bipush            
   0406 ➢ bipush            
   0408 ➢ castore           
   0409 ➢ dup               
   0410 ➢ bipush            
   0412 ➢ bipush            
   0414 ➢ castore           
   0415 ➢ dup               
   0416 ➢ bipush            
   0418 ➢ bipush            
   0420 ➢ castore           
   0421 ➢ dup               
   0422 ➢ bipush            
   0424 ➢ bipush            
   0426 ➢ castore           
   0427 ➢ dup               
   0428 ➢ bipush            
   0430 ➢ bipush            
   0432 ➢ castore           
   0433 ➢ dup               
   0434 ➢ bipush            
   0436 ➢ bipush            
   0438 ➢ castore           
   0439 ➢ dup               
   0440 ➢ bipush            
   0442 ➢ bipush            
   0444 ➢ castore           
   0445 ➢ dup               
   0446 ➢ bipush            
   0448 ➢ bipush            
   0450 ➢ castore           
   0451 ➢ dup               
   0452 ➢ bipush            
   0454 ➢ bipush            
   0456 ➢ castore           
   0457 ➢ dup               
   0458 ➢ bipush            
   0460 ➢ bipush            
   0462 ➢ castore           
   0463 ➢ dup               
   0464 ➢ bipush            
   0466 ➢ bipush            
   0468 ➢ castore           
   0469 ➢ dup               
   0470 ➢ bipush            
   0472 ➢ bipush            
   0474 ➢ castore           
   0475 ➢ dup               
   0476 ➢ bipush            
   0478 ➢ bipush            
   0480 ➢ castore           
   0481 ➢ dup               
   0482 ➢ bipush            
   0484 ➢ bipush            
   0486 ➢ castore           
   0487 ➢ dup               
   0488 ➢ bipush            
   0490 ➢ bipush            
   0492 ➢ castore           
   0493 ➢ dup               
   0494 ➢ bipush            
   0496 ➢ bipush            
   0498 ➢ castore           
   0499 ➢ dup               
   0500 ➢ bipush            
   0502 ➢ bipush            
   0504 ➢ castore           
   0505 ➢ dup               
   0506 ➢ bipush            
   0508 ➢ bipush            
   0510 ➢ castore           
   0511 ➢ dup               
   0512 ➢ bipush            
   0514 ➢ bipush            
   0516 ➢ castore           
   0517 ➢ dup               
   0518 ➢ bipush            
   0520 ➢ bipush            
   0522 ➢ castore           
   0523 ➢ dup               
   0524 ➢ bipush            
   0526 ➢ bipush            
   0528 ➢ castore           
   0529 ➢ dup               
   0530 ➢ bipush            
   0532 ➢ bipush            
   0534 ➢ castore           
   0535 ➢ dup               
   0536 ➢ bipush            
   0538 ➢ bipush            
   0540 ➢ castore           
   0541 ➢ dup               
   0542 ➢ bipush            
   0544 ➢ bipush            
   0546 ➢ castore           
   0547 ➢ dup               
   0548 ➢ bipush            
   0550 ➢ bipush            
   0552 ➢ castore           
   0553 ➢ dup               
   0554 ➢ bipush            
   0556 ➢ bipush            
   0558 ➢ castore           
   0559 ➢ dup               
   0560 ➢ bipush            
   0562 ➢ bipush            
   0564 ➢ castore           
   0565 ➢ dup               
   0566 ➢ bipush            
   0568 ➢ bipush            
   0570 ➢ castore           
   0571 ➢ dup               
   0572 ➢ bipush            
   0574 ➢ bipush            
   0576 ➢ castore           
   0577 ➢ dup               
   0578 ➢ bipush            
   0580 ➢ bipush            
   0582 ➢ castore           
   0583 ➢ dup               
   0584 ➢ bipush            
   0586 ➢ bipush            
   0588 ➢ castore           
   0589 ➢ dup               
   0590 ➢ bipush            
   0592 ➢ bipush            
   0594 ➢ castore           
   0595 ➢ dup               
   0596 ➢ bipush            
   0598 ➢ bipush            
   0600 ➢ castore           
   0601 ➢ dup               
   0602 ➢ bipush            
   0604 ➢ bipush            
   0606 ➢ castore           
   0607 ➢ dup               
   0608 ➢ bipush            
   0610 ➢ bipush            
   0612 ➢ castore           
   0613 ➢ dup               
   0614 ➢ bipush            
   0616 ➢ bipush            
   0618 ➢ castore           
   0619 ➢ dup               
   0620 ➢ bipush            
   0622 ➢ bipush            
   0624 ➢ castore           
   0625 ➢ dup               
   0626 ➢ bipush            
   0628 ➢ bipush            
   0630 ➢ castore           
   0631 ➢ dup               
   0632 ➢ bipush            
   0634 ➢ bipush            
   0636 ➢ castore           
   0637 ➢ dup               
   0638 ➢ bipush            
   0640 ➢ bipush            
   0642 ➢ castore           
   0643 ➢ dup               
   0644 ➢ bipush            
   0646 ➢ bipush            
   0648 ➢ castore           
   0649 ➢ dup               
   0650 ➢ bipush            
   0652 ➢ bipush            
   0654 ➢ castore           
   0655 ➢ dup               
   0656 ➢ bipush            
   0658 ➢ bipush            
   0660 ➢ castore           
   0661 ➢ dup               
   0662 ➢ bipush            
   0664 ➢ bipush            
   0666 ➢ castore           
   0667 ➢ dup               
   0668 ➢ bipush            
   0670 ➢ bipush            
   0672 ➢ castore           
   0673 ➢ dup               
   0674 ➢ bipush            
   0676 ➢ bipush            
   0678 ➢ castore           
   0679 ➢ dup               
   0680 ➢ bipush            
   0682 ➢ bipush            
   0684 ➢ castore           
   0685 ➢ dup               
   0686 ➢ bipush            
   0688 ➢ bipush            
   0690 ➢ castore           
   0691 ➢ dup               
   0692 ➢ bipush            
   0694 ➢ bipush            
   0696 ➢ castore           
   0697 ➢ dup               
   0698 ➢ bipush            
   0700 ➢ bipush            
   0702 ➢ castore           
   0703 ➢ dup               
   0704 ➢ bipush            
   0706 ➢ bipush            
   0708 ➢ castore           
   0709 ➢ dup               
   0710 ➢ bipush            
   0712 ➢ bipush            
   0714 ➢ castore           
   0715 ➢ dup               
   0716 ➢ bipush            
   0718 ➢ bipush            
   0720 ➢ castore           
   0721 ➢ dup               
   0722 ➢ bipush            
   0724 ➢ bipush            
   0726 ➢ castore           
   0727 ➢ dup               
   0728 ➢ bipush            
   0730 ➢ bipush            
   0732 ➢ castore           
   0733 ➢ dup               
   0734 ➢ bipush            
   0736 ➢ bipush            
   0738 ➢ castore           
   0739 ➢ dup               
   0740 ➢ bipush            
   0742 ➢ bipush            
   0744 ➢ castore           
   0745 ➢ dup               
   0746 ➢ bipush            
   0748 ➢ bipush            
   0750 ➢ castore           
   0751 ➢ dup               
   0752 ➢ bipush            
   0754 ➢ bipush            
   0756 ➢ castore           
   0757 ➢ dup               
   0758 ➢ bipush            
   0760 ➢ bipush            
   0762 ➢ castore           
   0763 ➢ dup               
   0764 ➢ bipush            
   0766 ➢ bipush            
   0768 ➢ castore           
   0769 ➢ dup               
   0770 ➢ bipush            
   0772 ➢ bipush            
   0774 ➢ castore           
   0775 ➢ dup               
   0776 ➢ bipush            
   0778 ➢ bipush            
   0780 ➢ castore           
   0781 ➢ dup               
   0782 ➢ bipush            
   0784 ➢ bipush            
   0786 ➢ castore           
   0787 ➢ dup               
   0788 ➢ bipush            
   0790 ➢ bipush            
   0792 ➢ castore           
   0793 ➢ dup               
   0794 ➢ bipush            
   0796 ➢ bipush            
   0798 ➢ castore           
   0799 ➢ dup               
   0800 ➢ bipush            
   0802 ➢ bipush            
   0804 ➢ castore           
   0805 ➢ dup               
   0806 ➢ bipush            
   0808 ➢ bipush            
   0810 ➢ castore           
   0811 ➢ dup               
   0812 ➢ bipush            
   0814 ➢ bipush            
   0816 ➢ castore           
   0817 ➢ dup               
   0818 ➢ bipush            
   0820 ➢ bipush            
   0822 ➢ castore           
   0823 ➢ putstatic          #22
   0826 ➢ bipush            
   0828 ➢ newarray          
   0830 ➢ dup               
   0831 ➢ iconst_0          
   0832 ➢ bipush            
   0834 ➢ castore           
   0835 ➢ dup               
   0836 ➢ iconst_1          
   0837 ➢ bipush            
   0839 ➢ castore           
   0840 ➢ dup               
   0841 ➢ iconst_2          
   0842 ➢ bipush            
   0844 ➢ castore           
   0845 ➢ dup               
   0846 ➢ iconst_3          
   0847 ➢ bipush            
   0849 ➢ castore           
   0850 ➢ dup               
   0851 ➢ iconst_4          
   0852 ➢ bipush            
   0854 ➢ castore           
   0855 ➢ dup               
   0856 ➢ iconst_5          
   0857 ➢ bipush            
   0859 ➢ castore           
   0860 ➢ dup               
   0861 ➢ bipush            
   0863 ➢ bipush            
   0865 ➢ castore           
   0866 ➢ dup               
   0867 ➢ bipush            
   0869 ➢ bipush            
   0871 ➢ castore           
   0872 ➢ dup               
   0873 ➢ bipush            
   0875 ➢ bipush            
   0877 ➢ castore           
   0878 ➢ dup               
   0879 ➢ bipush            
   0881 ➢ bipush            
   0883 ➢ castore           
   0884 ➢ dup               
   0885 ➢ bipush            
   0887 ➢ bipush            
   0889 ➢ castore           
   0890 ➢ dup               
   0891 ➢ bipush            
   0893 ➢ bipush            
   0895 ➢ castore           
   0896 ➢ dup               
   0897 ➢ bipush            
   0899 ➢ bipush            
   0901 ➢ castore           
   0902 ➢ dup               
   0903 ➢ bipush            
   0905 ➢ bipush            
   0907 ➢ castore           
   0908 ➢ dup               
   0909 ➢ bipush            
   0911 ➢ bipush            
   0913 ➢ castore           
   0914 ➢ dup               
   0915 ➢ bipush            
   0917 ➢ bipush            
   0919 ➢ castore           
   0920 ➢ dup               
   0921 ➢ bipush            
   0923 ➢ bipush            
   0925 ➢ castore           
   0926 ➢ dup               
   0927 ➢ bipush            
   0929 ➢ bipush            
   0931 ➢ castore           
   0932 ➢ dup               
   0933 ➢ bipush            
   0935 ➢ bipush            
   0937 ➢ castore           
   0938 ➢ dup               
   0939 ➢ bipush            
   0941 ➢ bipush            
   0943 ➢ castore           
   0944 ➢ dup               
   0945 ➢ bipush            
   0947 ➢ bipush            
   0949 ➢ castore           
   0950 ➢ dup               
   0951 ➢ bipush            
   0953 ➢ bipush            
   0955 ➢ castore           
   0956 ➢ dup               
   0957 ➢ bipush            
   0959 ➢ bipush            
   0961 ➢ castore           
   0962 ➢ dup               
   0963 ➢ bipush            
   0965 ➢ bipush            
   0967 ➢ castore           
   0968 ➢ dup               
   0969 ➢ bipush            
   0971 ➢ bipush            
   0973 ➢ castore           
   0974 ➢ dup               
   0975 ➢ bipush            
   0977 ➢ bipush            
   0979 ➢ castore           
   0980 ➢ dup               
   0981 ➢ bipush            
   0983 ➢ bipush            
   0985 ➢ castore           
   0986 ➢ dup               
   0987 ➢ bipush            
   0989 ➢ bipush            
   0991 ➢ castore           
   0992 ➢ dup               
   0993 ➢ bipush            
   0995 ➢ bipush            
   0997 ➢ castore           
   0998 ➢ dup               
   0999 ➢ bipush            
   1001 ➢ bipush            
   1003 ➢ castore           
   1004 ➢ dup               
   1005 ➢ bipush            
   1007 ➢ bipush            
   1009 ➢ castore           
   1010 ➢ dup               
   1011 ➢ bipush            
   1013 ➢ bipush            
   1015 ➢ castore           
   1016 ➢ dup               
   1017 ➢ bipush            
   1019 ➢ bipush            
   1021 ➢ castore           
   1022 ➢ dup               
   1023 ➢ bipush            
   1025 ➢ bipush            
   1027 ➢ castore           
   1028 ➢ dup               
   1029 ➢ bipush            
   1031 ➢ bipush            
   1033 ➢ castore           
   1034 ➢ dup               
   1035 ➢ bipush            
   1037 ➢ bipush            
   1039 ➢ castore           
   1040 ➢ dup               
   1041 ➢ bipush            
   1043 ➢ bipush            
   1045 ➢ castore           
   1046 ➢ dup               
   1047 ➢ bipush            
   1049 ➢ bipush            
   1051 ➢ castore           
   1052 ➢ dup               
   1053 ➢ bipush            
   1055 ➢ bipush            
   1057 ➢ castore           
   1058 ➢ dup               
   1059 ➢ bipush            
   1061 ➢ bipush            
   1063 ➢ castore           
   1064 ➢ dup               
   1065 ➢ bipush            
   1067 ➢ bipush            
   1069 ➢ castore           
   1070 ➢ dup               
   1071 ➢ bipush            
   1073 ➢ bipush            
   1075 ➢ castore           
   1076 ➢ dup               
   1077 ➢ bipush            
   1079 ➢ bipush            
   1081 ➢ castore           
   1082 ➢ dup               
   1083 ➢ bipush            
   1085 ➢ bipush            
   1087 ➢ castore           
   1088 ➢ dup               
   1089 ➢ bipush            
   1091 ➢ bipush            
   1093 ➢ castore           
   1094 ➢ dup               
   1095 ➢ bipush            
   1097 ➢ bipush            
   1099 ➢ castore           
   1100 ➢ dup               
   1101 ➢ bipush            
   1103 ➢ bipush            
   1105 ➢ castore           
   1106 ➢ dup               
   1107 ➢ bipush            
   1109 ➢ bipush            
   1111 ➢ castore           
   1112 ➢ dup               
   1113 ➢ bipush            
   1115 ➢ bipush            
   1117 ➢ castore           
   1118 ➢ dup               
   1119 ➢ bipush            
   1121 ➢ bipush            
   1123 ➢ castore           
   1124 ➢ dup               
   1125 ➢ bipush            
   1127 ➢ bipush            
   1129 ➢ castore           
   1130 ➢ dup               
   1131 ➢ bipush            
   1133 ➢ bipush            
   1135 ➢ castore           
   1136 ➢ dup               
   1137 ➢ bipush            
   1139 ➢ bipush            
   1141 ➢ castore           
   1142 ➢ dup               
   1143 ➢ bipush            
   1145 ➢ bipush            
   1147 ➢ castore           
   1148 ➢ dup               
   1149 ➢ bipush            
   1151 ➢ bipush            
   1153 ➢ castore           
   1154 ➢ dup               
   1155 ➢ bipush            
   1157 ➢ bipush            
   1159 ➢ castore           
   1160 ➢ dup               
   1161 ➢ bipush            
   1163 ➢ bipush            
   1165 ➢ castore           
   1166 ➢ dup               
   1167 ➢ bipush            
   1169 ➢ bipush            
   1171 ➢ castore           
   1172 ➢ dup               
   1173 ➢ bipush            
   1175 ➢ bipush            
   1177 ➢ castore           
   1178 ➢ dup               
   1179 ➢ bipush            
   1181 ➢ bipush            
   1183 ➢ castore           
   1184 ➢ dup               
   1185 ➢ bipush            
   1187 ➢ bipush            
   1189 ➢ castore           
   1190 ➢ dup               
   1191 ➢ bipush            
   1193 ➢ bipush            
   1195 ➢ castore           
   1196 ➢ dup               
   1197 ➢ bipush            
   1199 ➢ bipush            
   1201 ➢ castore           
   1202 ➢ dup               
   1203 ➢ bipush            
   1205 ➢ bipush            
   1207 ➢ castore           
   1208 ➢ dup               
   1209 ➢ bipush            
   1211 ➢ bipush            
   1213 ➢ castore           
   1214 ➢ dup               
   1215 ➢ bipush            
   1217 ➢ bipush            
   1219 ➢ castore           
   1220 ➢ dup               
   1221 ➢ bipush            
   1223 ➢ bipush            
   1225 ➢ castore           
   1226 ➢ dup               
   1227 ➢ bipush            
   1229 ➢ bipush            
   1231 ➢ castore           
   1232 ➢ dup               
   1233 ➢ bipush            
   1235 ➢ bipush            
   1237 ➢ castore           
   1238 ➢ dup               
   1239 ➢ bipush            
   1241 ➢ bipush            
   1243 ➢ castore           
   1244 ➢ dup               
   1245 ➢ bipush            
   1247 ➢ bipush            
   1249 ➢ castore           
   1250 ➢ dup               
   1251 ➢ bipush            
   1253 ➢ bipush            
   1255 ➢ castore           
   1256 ➢ dup               
   1257 ➢ bipush            
   1259 ➢ bipush            
   1261 ➢ castore           
   1262 ➢ dup               
   1263 ➢ bipush            
   1265 ➢ bipush            
   1267 ➢ castore           
   1268 ➢ dup               
   1269 ➢ bipush            
   1271 ➢ bipush            
   1273 ➢ castore           
   1274 ➢ dup               
   1275 ➢ bipush            
   1277 ➢ bipush            
   1279 ➢ castore           
   1280 ➢ dup               
   1281 ➢ bipush            
   1283 ➢ bipush            
   1285 ➢ castore           
   1286 ➢ dup               
   1287 ➢ bipush            
   1289 ➢ bipush            
   1291 ➢ castore           
   1292 ➢ dup               
   1293 ➢ bipush            
   1295 ➢ bipush            
   1297 ➢ castore           
   1298 ➢ dup               
   1299 ➢ bipush            
   1301 ➢ bipush            
   1303 ➢ castore           
   1304 ➢ dup               
   1305 ➢ bipush            
   1307 ➢ bipush            
   1309 ➢ castore           
   1310 ➢ dup               
   1311 ➢ bipush            
   1313 ➢ bipush            
   1315 ➢ castore           
   1316 ➢ dup               
   1317 ➢ bipush            
   1319 ➢ bipush            
   1321 ➢ castore           
   1322 ➢ dup               
   1323 ➢ bipush            
   1325 ➢ bipush            
   1327 ➢ castore           
   1328 ➢ dup               
   1329 ➢ bipush            
   1331 ➢ bipush            
   1333 ➢ castore           
   1334 ➢ dup               
   1335 ➢ bipush            
   1337 ➢ bipush            
   1339 ➢ castore           
   1340 ➢ dup               
   1341 ➢ bipush            
   1343 ➢ bipush            
   1345 ➢ castore           
   1346 ➢ dup               
   1347 ➢ bipush            
   1349 ➢ bipush            
   1351 ➢ castore           
   1352 ➢ dup               
   1353 ➢ bipush            
   1355 ➢ bipush            
   1357 ➢ castore           
   1358 ➢ dup               
   1359 ➢ bipush            
   1361 ➢ bipush            
   1363 ➢ castore           
   1364 ➢ dup               
   1365 ➢ bipush            
   1367 ➢ bipush            
   1369 ➢ castore           
   1370 ➢ dup               
   1371 ➢ bipush            
   1373 ➢ bipush            
   1375 ➢ castore           
   1376 ➢ dup               
   1377 ➢ bipush            
   1379 ➢ bipush            
   1381 ➢ castore           
   1382 ➢ dup               
   1383 ➢ bipush            
   1385 ➢ bipush            
   1387 ➢ castore           
   1388 ➢ dup               
   1389 ➢ bipush            
   1391 ➢ bipush            
   1393 ➢ castore           
   1394 ➢ dup               
   1395 ➢ bipush            
   1397 ➢ bipush            
   1399 ➢ castore           
   1400 ➢ dup               
   1401 ➢ bipush            
   1403 ➢ bipush            
   1405 ➢ castore           
   1406 ➢ dup               
   1407 ➢ bipush            
   1409 ➢ bipush            
   1411 ➢ castore           
   1412 ➢ dup               
   1413 ➢ bipush            
   1415 ➢ bipush            
   1417 ➢ castore           
   1418 ➢ dup               
   1419 ➢ bipush            
   1421 ➢ bipush            
   1423 ➢ castore           
   1424 ➢ putstatic          #21
   1427 ➢ bipush            
   1429 ➢ newarray          
   1431 ➢ dup               
   1432 ➢ iconst_0          
   1433 ➢ bipush            
   1435 ➢ iastore           
   1436 ➢ dup               
   1437 ➢ iconst_1          
   1438 ➢ bipush            
   1440 ➢ iastore           
   1441 ➢ dup               
   1442 ➢ iconst_2          
   1443 ➢ sipush            
   1446 ➢ iastore           
   1447 ➢ dup               
   1448 ➢ iconst_3          
   1449 ➢ sipush            
   1452 ➢ iastore           
   1453 ➢ dup               
   1454 ➢ iconst_4          
   1455 ➢ ldc                #88
   1457 ➢ iastore           
   1458 ➢ dup               
   1459 ➢ iconst_5          
   1460 ➢ ldc                #89
   1462 ➢ iastore           
   1463 ➢ dup               
   1464 ➢ bipush            
   1466 ➢ ldc                #90
   1468 ➢ iastore           
   1469 ➢ dup               
   1470 ➢ bipush            
   1472 ➢ ldc                #91
   1474 ➢ iastore           
   1475 ➢ dup               
   1476 ➢ bipush            
   1478 ➢ ldc                #92
   1480 ➢ iastore           
   1481 ➢ dup               
   1482 ➢ bipush            
   1484 ➢ ldc                #93
   1486 ➢ iastore           
   1487 ➢ putstatic          #24
   1490 ➢ return            
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0026 ➢ invokestatic       #8
🍷 java/lang/Integer.toHexString(I)Ljava/lang/String; 
   0000 ➢ iload_0           
   0001 ➢ iconst_4          
   0002 ➢ invokestatic       #10
🍷 java/lang/Integer.toUnsignedString0(II)Ljava/lang/String; 
   0000 ➢ bipush            
   0002 ➢ iload_0           
   0003 ➢ invokestatic       #11
🍷 java/lang/Integer.numberOfLeadingZeros(I)I 
   0000 ➢ iload_0           
   0001 ➢ ifne              
   0007 ➢ iconst_1          
   0008 ➢ istore_1          
   0009 ➢ iload_0           
   0010 ➢ bipush            
   0012 ➢ iushr             
   0013 ➢ ifne              
   0016 ➢ iinc               #1
   0019 ➢ iload_0           
   0020 ➢ bipush            
   0022 ➢ ishl              
   0023 ➢ istore_0          
   0024 ➢ iload_0           
   0025 ➢ bipush            
   0027 ➢ iushr             
   0028 ➢ ifne              
   0031 ➢ iinc               #1
   0034 ➢ iload_0           
   0035 ➢ bipush            
   0037 ➢ ishl              
   0038 ➢ istore_0          
   0039 ➢ iload_0           
   0040 ➢ bipush            
   0042 ➢ iushr             
   0043 ➢ ifne              
   0046 ➢ iinc               #1
   0049 ➢ iload_0           
   0050 ➢ iconst_4          
   0051 ➢ ishl              
   0052 ➢ istore_0          
   0053 ➢ iload_0           
   0054 ➢ bipush            
   0056 ➢ iushr             
   0057 ➢ ifne              
   0060 ➢ iinc               #1
   0063 ➢ iload_0           
   0064 ➢ iconst_2          
   0065 ➢ ishl              
   0066 ➢ istore_0          
   0067 ➢ iload_1           
   0068 ➢ iload_0           
   0069 ➢ bipush            
   0071 ➢ iushr             
   0072 ➢ isub              
   0073 ➢ istore_1          
   0074 ➢ iload_1           
   0075 ➢ ireturn           
🍷 java/lang/Integer.toUnsignedString0(II)Ljava/lang/String; 
   0006 ➢ isub              
   0007 ➢ istore_2          
   0008 ➢ iload_2           
   0009 ➢ iload_1           
   0010 ➢ iconst_1          
   0011 ➢ isub              
   0012 ➢ iadd              
   0013 ➢ iload_1           
   0014 ➢ idiv              
   0015 ➢ iconst_1          
   0016 ➢ invokestatic       #12
🍷 java/lang/Math.max(II)I 
   0000 ➢ iload_0           
   0001 ➢ iload_1           
   0002 ➢ if_icmplt         
   0005 ➢ iload_0           
   0006 ➢ goto              
   0010 ➢ ireturn           
🍷 java/lang/Integer.toUnsignedString0(II)Ljava/lang/String; 
   0019 ➢ istore_3          
   0020 ➢ iload_3           
   0021 ➢ newarray          
   0023 ➢ astore             #4
   0025 ➢ iload_0           
   0026 ➢ iload_1           
   0027 ➢ aload              #4
   0029 ➢ iconst_0          
   0030 ➢ iload_3           
   0031 ➢ invokestatic       #13
🍷 java/lang/Integer.formatUnsignedInt(II[CII)I 
   0000 ➢ iload              #4
   0002 ➢ istore             #5
   0004 ➢ iconst_1          
   0005 ➢ iload_1           
   0006 ➢ ishl              
   0007 ➢ istore             #6
   0009 ➢ iload              #6
   0011 ➢ iconst_1          
   0012 ➢ isub              
   0013 ➢ istore             #7
   0015 ➢ aload_2           
   0016 ➢ iload_3           
   0017 ➢ iinc               #5
   0020 ➢ iload              #5
   0022 ➢ iadd              
   0023 ➢ getstatic          #5
   0026 ➢ iload_0           
   0027 ➢ iload              #7
   0029 ➢ iand              
   0030 ➢ caload            
   0031 ➢ castore           
   0032 ➢ iload_0           
   0033 ➢ iload_1           
   0034 ➢ iushr             
   0035 ➢ istore_0          
   0036 ➢ iload_0           
   0037 ➢ ifeq              
   0045 ➢ iload              #5
   0047 ➢ ireturn           
🍷 java/lang/Integer.toUnsignedString0(II)Ljava/lang/String; 
   0034 ➢ pop               
   0035 ➢ new                #6
   0038 ➢ dup               
   0039 ➢ aload              #4
   0041 ➢ iconst_1          
   0042 ➢ invokespecial      #14
🍷 java/lang/String.<init>([CZ)V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 java/lang/String.<init>([CZ)V 
   0004 ➢ aload_0           
   0005 ➢ aload_1           
   0006 ➢ putfield           #3
   0009 ➢ return            
🍷 java/lang/Integer.toUnsignedString0(II)Ljava/lang/String; 
   0045 ➢ areturn           
🍷 java/lang/Integer.toHexString(I)Ljava/lang/String; 
   0005 ➢ areturn           
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0029 ➢ invokevirtual      #5
🍷 java/lang/StringBuilder.append(Ljava/lang/String;)Ljava/lang/StringBuilder; 
   0000 ➢ aload_0           
   0001 ➢ aload_1           
   0002 ➢ invokespecial      #8
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0000 ➢ aload_1           
   0001 ➢ ifnonnull         
   0009 ➢ aload_1           
   0010 ➢ invokevirtual      #29
🍷 java/lang/String.length()I 
   0000 ➢ aload_0           
   0001 ➢ getfield           #3
   0004 ➢ arraylength       
   0005 ➢ ireturn           
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0013 ➢ istore_2          
   0014 ➢ aload_0           
   0015 ➢ aload_0           
   0016 ➢ getfield           #4
   0019 ➢ iload_2           
   0020 ➢ iadd              
   0021 ➢ invokespecial      #5
🍷 java/lang/AbstractStringBuilder.ensureCapacityInternal(I)V 
   0000 ➢ iload_1           
   0001 ➢ aload_0           
   0002 ➢ getfield           #3
   0005 ➢ arraylength       
   0006 ➢ isub              
   0007 ➢ ifle              
   0026 ➢ return            
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0024 ➢ aload_1           
   0025 ➢ iconst_0          
   0026 ➢ iload_2           
   0027 ➢ aload_0           
   0028 ➢ getfield           #3
   0031 ➢ aload_0           
   0032 ➢ getfield           #4
   0035 ➢ invokevirtual      #30
🍷 java/lang/String.getChars(II[CI)V 
   0000 ➢ iload_1           
   0001 ➢ ifge              
   0013 ➢ iload_2           
   0014 ➢ aload_0           
   0015 ➢ getfield           #3
   0018 ➢ arraylength       
   0019 ➢ if_icmple         
   0031 ➢ iload_1           
   0032 ➢ iload_2           
   0033 ➢ if_icmple         
   0047 ➢ aload_0           
   0048 ➢ getfield           #3
   0051 ➢ iload_1           
   0052 ➢ aload_3           
   0053 ➢ iload              #4
   0055 ➢ iload_2           
   0056 ➢ iload_1           
   0057 ➢ isub              
   0058 ➢ invokestatic       #37
🍺invoke native method java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V 😎

   0061 ➢ return            
🍷 java/lang/AbstractStringBuilder.append(Ljava/lang/String;)Ljava/lang/AbstractStringBuilder; 
   0038 ➢ aload_0           
   0039 ➢ dup               
   0040 ➢ getfield           #4
   0043 ➢ iload_2           
   0044 ➢ iadd              
   0045 ➢ putfield           #4
   0048 ➢ aload_0           
   0049 ➢ areturn           
🍷 java/lang/StringBuilder.append(Ljava/lang/String;)Ljava/lang/StringBuilder; 
   0005 ➢ pop               
   0006 ➢ aload_0           
   0007 ➢ areturn           
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0032 ➢ invokevirtual      #9
🍷 java/lang/StringBuilder.toString()Ljava/lang/String; 
   0000 ➢ new                #41
   0003 ➢ dup               
   0004 ➢ aload_0           
   0005 ➢ getfield           #42
   0008 ➢ iconst_0          
   0009 ➢ aload_0           
   0010 ➢ getfield           #43
   0013 ➢ invokespecial      #44
🍷 java/lang/String.<init>([CII)V 
   0000 ➢ aload_0           
   0001 ➢ invokespecial      #1
🍷 java/lang/Object.<init>()V 
   0000 ➢ return            
🍷 java/lang/String.<init>([CII)V 
   0004 ➢ iload_2           
   0005 ➢ ifge              
   0017 ➢ iload_3           
   0018 ➢ ifgt              
   0050 ➢ iload_2           
   0051 ➢ aload_1           
   0052 ➢ arraylength       
   0053 ➢ iload_3           
   0054 ➢ isub              
   0055 ➢ if_icmple         
   0069 ➢ aload_0           
   0070 ➢ aload_1           
   0071 ➢ iload_2           
   0072 ➢ iload_2           
   0073 ➢ iload_3           
   0074 ➢ iadd              
   0075 ➢ invokestatic       #8
🍷 java/util/Arrays.copyOfRange([CII)[C 
   0000 ➢ iload_2           
   0001 ➢ iload_1           
   0002 ➢ isub              
   0003 ➢ istore_3          
   0004 ➢ iload_3           
   0005 ➢ ifge              
   0039 ➢ iload_3           
   0040 ➢ newarray          
   0042 ➢ astore             #4
   0044 ➢ aload_0           
   0045 ➢ iload_1           
   0046 ➢ aload              #4
   0048 ➢ iconst_0          
   0049 ➢ aload_0           
   0050 ➢ arraylength       
   0051 ➢ iload_1           
   0052 ➢ isub              
   0053 ➢ iload_3           
   0054 ➢ invokestatic       #100
🍷 java/lang/Math.min(II)I 
   0000 ➢ iload_0           
   0001 ➢ iload_1           
   0002 ➢ if_icmpgt         
   0009 ➢ iload_1           
   0010 ➢ ireturn           
🍷 java/util/Arrays.copyOfRange([CII)[C 
   0057 ➢ invokestatic       #65
🍺invoke native method java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V 😎

   0060 ➢ aload              #4
   0062 ➢ areturn           
🍷 java/lang/String.<init>([CII)V 
   0078 ➢ putfield           #3
   0081 ➢ return            
🍷 java/lang/StringBuilder.toString()Ljava/lang/String; 
   0016 ➢ areturn           
🍷 java/lang/Object.toString()Ljava/lang/String; 
   0035 ➢ areturn           
🍷 C.ss()Ljava/lang/String; 
   0012 ➢ areturn           
🍷 Main.main([Ljava/lang/String;)V 
   0109 ➢ invokevirtual      #15
🍷 GVM.println(Ljava/lang/Object;)V 
   0000 ➢ aload_0           
   0001 ➢ aload_1           
   0002 ➢ invokestatic       #2
🍷 java/lang/String.valueOf(Ljava/lang/Object;)Ljava/lang/String; 
   0000 ➢ aload_0           
   0001 ➢ ifnonnull         
   0009 ➢ aload_0           
   0010 ➢ invokevirtual      #135
🍷 java/lang/String.toString()Ljava/lang/String; 
   0000 ➢ aload_0           
   0001 ➢ areturn           
🍷 java/lang/String.valueOf(Ljava/lang/Object;)Ljava/lang/String; 
   0013 ➢ areturn           
🍷 GVM.println(Ljava/lang/Object;)V 
   0005 ➢ invokespecial      #3
🍺invoke native method GVM.println(Ljava/lang/String;)V 😎

   0008 ➢ return            
🍷 Main.main([Ljava/lang/String;)V 
   0112 ➢ return            
[I@1   
```