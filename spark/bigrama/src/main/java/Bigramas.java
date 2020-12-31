import java.util.ArrayList;
import java.util.List;
import java.util.regex.Pattern;
import java.io.FileWriter;

import scala.Tuple2;

import com.google.common.collect.Iterables;

import org.apache.spark.api.java.JavaPairRDD;
import org.apache.spark.api.java.JavaRDD;
import org.apache.spark.api.java.function.Function2;
import org.apache.spark.sql.SparkSession;

public final class Bigramas{
	
	private static final Pattern SPACES = Pattern.compile("\\s+");
	
	public static void main(String[] args) throws Exception {
		
		SparkSession spark = SparkSession
		  .builder()
		  .appName("Bigramas")
		  .getOrCreate();
		  
		JavaRDD<String> lines = spark.read().textFile(args[0]).javaRDD();
		
		StringTokenizer itr = new StringTokenizer(lines.toString());
		
		while (itr.hasMoreTokens()) {
		//se tokeniza y se arma el par <palabra1/palabra2, 1>
		//falta
		JavaPairRDD<String, Iterable<String>> bigramas = lines.mapToPair(s -> {
			String[] parts = SPACES.split(s);
			return new Tuple2<>(parts[0], parts[1]);//no esta bien tokenizado
		}).distinct().groupByKey().cache();
		}
		
		
		JavaPairRDD<String, Double> bigr = bigramas.reduceByKey(new Sum());

		
		
		List<Tuple2<String, Double>> output = bigr.collect();
		FileWriter out = new FileWriter(args[1]);
		for (Tuple2<?,?> tuple : output) {
		  out.write(tuple._1() + " / " + tuple._2() + ".\n");
		}
		out.close();
		
		spark.stop();
	}
}

