//// modules.php
<?hh
<<file:__EnableUnstableFeatures("modules")>>
new module a {}    // package pkg1
new module b.b1 {} // package pkg2 (include pkg1)

//// a.php
<?hh
<<file:__EnableUnstableFeatures("modules")>>

module a;
public class A {}

//// b.php
<?hh
<<file:__EnableUnstableFeatures("modules")>>

module b.b1;
public class B1 {}
public function test(): void {
   $a = new A(); // ok
}
