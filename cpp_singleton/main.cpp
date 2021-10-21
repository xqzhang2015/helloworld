#include <stdio.h>
#include <string>

class noncopyable
{
public:
    noncopyable(const noncopyable&) = delete;
    void operator=(const noncopyable&) = delete;
protected:
    noncopyable() = default;
    ~noncopyable() = default;
};

template <class T>
class Singleton : public noncopyable
{
public:
    static T* instance()
    {
        static T inst;
        return &inst;
    }

protected:
    Singleton() = default;
    ~Singleton() = default;
};

struct MyString : public Singleton<MyString>
{
    std::string str;
};

int main()
{
    printf("hello: %s\n", "world");

    MyString myStr;
    MyString::instance()->str = "I'm str member!";
    printf("myStr: %s\n", myStr.str.c_str());
    printf("MyString::instance: %s\n", MyString::instance()->str.c_str());

    return 0;
}

#if 0
 ⚙  helloworld/cpp_singleton   master ±✚  clang++ -std=c++11 main.cpp // -o singleton
 ⚙  helloworld/cpp_singleton   master ±✚  ./a.out
hello: world
myStr:
MyString::instance: I'm str member!
#endif