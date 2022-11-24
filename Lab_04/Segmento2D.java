import java.util.Scanner;
import java.util.ArrayList;

public class Segmento2D{
  //Overview: Classe che modella un segmento immutabile
  final Punto2D a,b;
  //costruttori

  //MODIFIES: this
  //EFFECTS: inizializzo un segmento con due punti non nulli
  //se i punti sono nulli o uguali lancia un IllegalArgumentException
  public Segmento2D(Punto2D a, Punto2D b) throws IllegalArgumentException{
    if (a==null){
      throw new IllegalArgumentException("Il punto a non può essere nullo");
    }
    if (b==null){
      throw new IllegalArgumentException("Il punto b non può essere nullo");
    }
    if (a.equals(b)){
      throw new IllegalArgumentException("Il segmento necessita di due punti diversi");
    }
    this.a=a;
    this.b=b;
  }

  public Punto2D getA(){
    return this.a;
  }

  public Punto2D getB(){
    return this.b;
  }

  //EFFECTS: recupera la langhezza del segmento
  public double length(){
    double c1=(this.getA().getX()-this.getB().getX());
    double c2=(this.getA().getY()-this.getB().getY());
    return Math.sqrt(c1*c1+c2*c2);
  }

  //Le invarianti sono che a e b
  //non devono essere uguali e non devono essere nulle
  private boolean repOk(){
    if (this.a==null){
      return false;
    }
    if (this.b==null){
      return false;
    }
    if (this.a.equals(b)){
      return false;
    }
    return true;
  }

  @Override
  public String toString(){
    return "segmento - a: (" + this.getA() + "), b: (" + this.getB() + "), lunghezza: "+this.length();
  }

  public static void main(String[] args){
    Scanner tastiera = new Scanner(System.in);
    ArrayList<Segmento2D> l = new ArrayList<Segmento2D>();
    double minlength = Double.parseDouble(args[0]);
    System.out.println("Inserisci i segmenti nel formato ax ay bx by (temina con CTRL+D)");
    while (tastiera.hasNext()){
      double ax = Double.parseDouble(tastiera.next());
      double ay = Double.parseDouble(tastiera.next());
      double bx = Double.parseDouble(tastiera.next());
      double by = Double.parseDouble(tastiera.next());
      l.add(new Segmento2D(new Punto2D(ax,ay),new Punto2D(bx,by)));
    }
    System.out.println("Segmenti di lunghezza superiore a " + minlength);
    for (Segmento2D segmento : l ) {
      if (segmento.length()> minlength){
        System.out.println(segmento);
      }
    }
    tastiera.close();
  }
}
