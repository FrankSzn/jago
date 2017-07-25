package jago

import "runtime"

func register_java_lang_System()  {
	register("java/lang/System.registerNatives()V", Java_java_lang_System_registerNatives)
	register("java/lang/System.setIn0(Ljava/io/InputStream;)V", Java_java_lang_System_setIn0)
	register("java/lang/System.setOut0(Ljava/io/PrintStream;)V", Java_java_lang_System_setOut0)
	register("java/lang/System.setErr0(Ljava/io/PrintStream;)V", Java_java_lang_System_setErr0)
	register("java/lang/System.currentTimeMillis()J", Java_java_lang_System_currentTimeMillis)
	register("java/lang/System.nanoTime()J", Java_java_lang_System_nanoTime)
	register("java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V", Java_java_lang_System_arraycopy)
	register("java/lang/System.identityHashCode(Ljava/lang/Object;)I", Java_java_lang_System_identityHashCode)
	register("java/lang/System.initProperties(Ljava/util/Properties;)Ljava/util/Properties;", Java_java_lang_System_initProperties)
}

// private static void registers()
func Java_java_lang_System_registerNatives() {}

// private static void setIn0(InputStream is)
func Java_java_lang_System_setIn0(is ObjectRef) {
	VM_setStaticVariable(VM_getClass("java/lang/System"), "in", "Ljava/io/InputStream;", is)
}

// private static void setOut0(PrintStream ps)
func Java_java_lang_System_setOut0(ps ObjectRef)  {
	VM_setStaticVariable(VM_getClass("java/lang/System"), "out", "Ljava/io/PrintStream;", ps)
}

// private static void setErr0(PrintStream ps)
func Java_java_lang_System_setErr0(ps ObjectRef)  {
	VM_setStaticVariable(VM_getClass("java/lang/System"), "err", "Ljava/io/PrintStream;", ps)
}

// public static long currentTimeMillis()
func Java_java_lang_System_currentTimeMillis() Long {
	return VM_currentTimeMillis()
}

// public static long nanoTime()
func Java_java_lang_System_nanoTime() Long {
	return VM_currentTimeNano()
}

// public static void arraycopy(Object fromArray, int fromIndex, Object toArray, int toIndex, int length)
func Java_java_lang_System_arraycopy(src Reference, srcPos Int, dest Reference, destPos Int, length Int) {
	if !src.Class().IsArray() || !dest.Class().IsArray() {
		Throw("java/lang/ArrayStoreException", "")
	}
	srcArrayRef := src.AsArrayRef()
	dstArrayRef := dest.AsArrayRef()

	if srcPos + length > srcArrayRef.Length() || destPos + length > dstArrayRef.Length() {
		Throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	for i := Int(0); i < length; i++ {
		dstArrayRef.SetElement(destPos + i, srcArrayRef.GetElement(srcPos + i))
	}
}

// public static int identityHashCode(Object object)
func Java_java_lang_System_identityHashCode(object Reference) Int {
	return VM_iHashCode(object)
}

// private static Properties initProperties(Properties properties)
func Java_java_lang_System_initProperties(properties ObjectRef) ObjectRef {

	//TODO
	//setProperty := properties.Class().GetMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	//VM_invokeMethod(THREAD_MANAGER.currentThread, setProperty, properties, NewJavaLangString("file.encoding"), NewJavaLangString("UTF-8"))
	//for key, val := range _sysProps {
	//	VM_invokeMethod(THREAD_MANAGER.currentThread, setProperty, properties, NewJavaLangString(key), NewJavaLangString(val))
	//}

	return properties
}

var _sysProps map[string]string = map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jago",
		"java.vendor.url":      "https://yangchao.me",
		"java.home":            "todo",
		"java.class.version":   "52.0",
		"java.class.path":      "todo",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
}